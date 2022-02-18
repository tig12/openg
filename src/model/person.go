/******************************************************************************
    Person

    @license    GPL
    @history    2021-07-14 15:48:55+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
    //"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
	"strconv"
)

// see init function at the end

type Person struct {
	Id             int
	Slug           string
	Sex            string
	Name           PersonName
	Occus          []string
	Birth          Event
	Death          Event
	Ids_in_sources map[string]string
	Ids_partial    interface{} // map[string]string
	Trust          interface{}
	Acts           []string
	History        []HistoryEntry
	Issues         []string
	Notes          []string
	// not stored in table person
	Groups []*PersonGroup
}

type PersonName struct {
	Usual     string
	Given     string
	Family    string
	Spouse    string
	Official  OfficialName
	Fame      FameName
	Nicknames []string
	Alter     []string
	Nobl      string
}

type OfficialName struct {
	Given  string
	Family string
}

type FameName struct {
	Full   string
	Given  string
	Family string
}

type HistoryEntry struct {
	Command string
	Date    string
	Source  string
	Raw     map[string]interface{}
	New     interface{}
}

// used by ajax
type AutocompletePerson struct {
    Slug string `json:"slug"`
    Day  string `json:"day"`
    Name string `json:"name"`
}



// Displayed names of the partial ids
var Ids_partial_labels = map[string]string{
	"lerrcp": "Gauquelin",
	"muller": "Müller",
	"cpara":  "Comité Para",
	"csicop": "CSICOP",
	"cfepp":  "CFEPP",
	"ertel":  "Ertel",
	"wd":     "Wikidata",
}

// ************************** PersonName *******************************

func (p *Person) String() string {
	return p.Slug
}

/**
    Returns a string representing a person name
**/
func (n *PersonName) DisplayedName() string {
	if n.Usual != "" {
		return n.Usual
	}
	if n.Fame.Full != "" {
		return n.Fame.Full
	}
	fam := n.Family
	if n.Nobl != "" {
		if string(n.Nobl[len(n.Nobl)-1]) == "'" {
			fam = n.Nobl + fam
		} else {
			fam = n.Nobl + " " + fam
		}
	}
	if fam != "" {
		if n.Fame.Given != "" {
			return fam + " " + n.Fame.Given
		}
	}
	if n.Given == "" {
		return fam
	}
	return fam + " " + n.Given
}

// ************************** Get one *******************************

/**
    Loads a person from database
**/
func GetPerson(restURL, slug string) (person *Person, err error) {
	url := restURL + "/person?slug=eq." + slug
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding person data "+slug)
	}
	persons := []Person{}
	if err = json.Unmarshal(responseData, &persons); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal person data "+slug+"\n"+string(responseData)+"\n")
	}
	if len(persons) > 1 {
		return nil, werr.New("Several persons with identical slug: " + slug)
	}
	if len(persons) == 0 {
		// do not return error because should be handled as 404 by caller
		return nil, nil
	}
	return &persons[0], nil
}

// ************************** Compute fields *******************************

/** Computes field 'Groups' of a person **/
func (p *Person) ComputeGroups(restURL string) (err error) {
	url := restURL + "/api_persongroop?person_id=eq." + strconv.Itoa(p.Id)
	response, err := http.Get(url)
	if err != nil {
		return werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return werr.Wrapf(err, "Error decoding PersonGroup data"+string(responseData)+"\n")
	}
	if err = json.Unmarshal(responseData, &p.Groups); err != nil {
		return werr.Wrapf(err, fmt.Sprintf("Error json Unmarshal PersonGroups\n%s\n", string(responseData)))
	}
	return nil
}


/** Computes field 'Acts' of a person **/
/** 
func (p *Person) ComputeActs(dirActs string) (err error) {
    file := dirActs
	//y, err := ioutil.ReadFile(dirActs)
	if err != nil {
	    if os.IsNotExist(err){
            return; // p.Acts remains empty
	    }
		panic(err)
	}
//	err = yaml.Unmarshal(y, &config)
	if err != nil {
		panic(err)
	}
}
**/

// ************************** Get many *******************************

// ==================== NOT USED - REMOVE ? ====================
func GetPersons(restURL string) (p []*Person, err error) {
	url := restURL + "/person?limit=10&offset=0"
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding persons data")
	}
	persons := []*Person{}
	if err = json.Unmarshal(responseData, &persons); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal persons data\n"+string(responseData)+"\n")
	}
	return persons, nil
}

/** 
    Used by ajax
**/
func GetPersonsAutocomplete(restURL, str string) (p []*AutocompletePerson, err error) {
	url := restURL + "/search?slug=like." + str + "*"
fmt.Println("=== model.GetPersonsAutocomplete url = " +url)
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding persons data")
	}
	if err = json.Unmarshal(responseData, &p); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal persons data\n"+string(responseData)+"\n")
	}
	return p, nil
}

// ************************** Get fields *******************************

/**
    Returns a person's birth date (day and time), format YYYY-MM-DD HH:MM:SS
    If Birth.Date exists, uses it.
    Otherwise uses field Birth.DateUT
**/
func (p *Person) GetBirthDate() string {
	return GetLegalOrUTDate(p.Birth.Date, p.Birth.DateUT)
}

/**
    Returns a person's death date (day and time), format YYYY-MM-DD HH:MM:SS
    If Death.Date exists, uses it.
    Otherwise uses field Death.DateUT
**/
func (p *Person) GetDeathDate() string {
	return GetLegalOrUTDate(p.Death.Date, p.Death.DateUT)
}

/**
    Returns a person's birth day, format YYYY-MM-DD
    If Birth.Date exists, uses it.
    Otherwise uses field Birth.DateUT
**/
func (p *Person) GetBirthDay() string {
	tmp := p.GetBirthDate()
	if len(tmp) == 0 {
		return ""
	}
	return tmp[:10]
}

/**
    Returns a person's death day, format YYYY-MM-DD
    If Death.Date exists, uses it.
    Otherwise uses field Death.DateUT
**/
func (p *Person) GetDeathDay() string {
	tmp := p.GetDeathDate()
	if len(tmp) == 0 {
		return ""
	}
	return tmp[:10]
}

/**
    Function used by Person and GroupMember
**/
func GetLegalOrUTDate(date, dateUT string) (res string) {
	if date != "" {
		res = date
	} else if dateUT != "" {
		res = dateUT
	} else {
		return ""
	}
	return res
}

func GetRawPersonSortedFields(source string) []string {
	switch source {
	case "a1", "a2", "a3", "a4", "a5", "a6":
		return RawPersonSortedFields["a"]
	case "d6":
		return RawPersonSortedFields["d6"]
	case "d10":
		return RawPersonSortedFields["d10"]
	case "e1", "e3":
		return RawPersonSortedFields["e"]
	case "afd1":
		return RawPersonSortedFields["afd1"]
	case "afd2":
		return RawPersonSortedFields["afd2"]
	case "afd1-100":
		return RawPersonSortedFields["afd1-100"]
	case "afd3":
		return RawPersonSortedFields["afd3"]
	case "afd5":
		return RawPersonSortedFields["afd5"]
	case "csi":
		return RawPersonSortedFields["csi"]
	default:
		return []string{}
	}
}

// ************************** Raw ordering *******************************
// Auxiliary structures to sort the raw fileds of a person

// map source slug => array of keys.
var RawPersonSortedFields map[string][]string

// Executed once at package loading
func init() {
	RawPersonSortedFields = map[string][]string{
		"a": {
			"YEA",
			"MON",
			"DAY",
			"PRO",
			"NUM",
			"COU",
			"H",
			"MN",
			"SEC",
			"TZ",
			"LAT",
			"LON",
			"COD",
			"CITY",
		},
		"d6": {
			"NUM",
			"DAY",
			"MON",
			"YEA",
			"H",
			"MN",
			"SEC",
			"LAT",
			"LON",
			"NAME",
		},
		"d10": {
			"NUM",
			"NAME",
			"PRO",
			"DAY",
			"MON",
			"YEA",
			"H",
			"TZ",
			"LAT",
			"LON",
			"CICO",
		},
		"e": {
			"NUM",
			"PRO",
			"NAME",
			"NOTE",
			"DAY",
			"MON",
			"YEA",
			"H",
			"CITY",
			"COD",
		},
		"afd5": {
			"NR",
			"SAMPLE",
			"GNR",
			"CODE",
			"NAME",
			"GEBDATUM",
			"JAHR",
			"GEBZEIT",
			"GEBORT",
			"LAENGE",
			"BREITE",
			"MODE",
			"KORR",
			"ELECTDAT",
			"ELECTAGE",
			"STBDATUM",
			"SONNE",
			"MOND",
			"VENUS",
			"MARS",
			"JUPITER",
			"SATURN",
			"SO_",
			"MO_",
			"VE_",
			"MA_",
			"JU_",
			"SA_",
			"PHAS_",
			"AUFAB",
			"NIENMO",
			"NIENVE",
			"NIENMA",
			"NIENJU",
			"NIENSA",
		},
		"afd1": {
			"NAME",
			"YEAR",
			"MONTH",
			"DAY",
			"HOUR",
			"MIN",
			"TZO",
			"PLACE",
			"LAT",
			"LG",
		},
		"afd1-100": {
			"MUID",
			"FNAME",
			"GNAME",
			"SEX",
			"DATE",
			"TZO",
			"PLACE",
			"C2",
			"LG",
			"LAT",
			"OCCU",
			"OPUS",
			"LEN",
		},
		"afd2": {
			"MUID",
			"NAME",
			"DATE",
			"TIME",
			"TZO",
			"TIMOD",
			"CY",
			"PLACE",
			"LAT",
			"LG",
			"OCCU",
			"BOOKS",
			"SOURCE",
			"GQ",
		},
		"afd3": {
			"MUID",
			"NAME",
			"DATE",
			"TIME",
			"TZO",
			"TIMOD",
			"CY",
			"PLACE",
			"LAT",
			"LG",
			"OCCU",
			"BOOKS",
			"SOURCE",
			"GQ",
		},
		"csi": {
			"Satz#",
			"NAME",
			"VORNAME",
			"GEBDAT",
			"GEBZEIT",
			"AMPM",
			"ZEITZONE",
			"GEBORT",
			"LO1",
			"LO2",
			"LA1",
			"LA2",
			"SPORTART",
			"MARS",
			"BATCH",
		},
	}
}
