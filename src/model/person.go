/******************************************************************************
    Person

    @license    GPL
    @history    2021-07-14 15:48:55+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"openg.local/openg/generic/wilk/werr"
	"fmt"
	"strconv"
	"strings"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// see init function at the end

type Person struct {
	// fields stored in table person
	Id    int
	Slug  string
	Sex   string
	Name  *PersonName
	Occus []string
	Birth *Event
	Death *Event
	// declared interface{} instead of map[string]string
	// because if empty, json has an empty array, not an empty map.
	// problem when building the json in php, cannot distinguish between an empty array and an empty map.
	Ids_in_sources interface{} // map[string]string
	Partial_ids    interface{} // map[string]string
	Trust          string
	Acts           *Acts `json:"acts"`
	History        []PersonHistoryEntry
	Issues         []*PersonIssue
	Notes          []string
	// not stored in table person
	Groups []*PersonGroup
}

/**
    Fields that can be present in an act, in field "person" or "extras".
    Put apart to avoid recursion problem
    // TODO : see if it can be used for struct Person
**/
type PartialPerson struct {
	// fields stored in table person
	Id             int
	Slug           string
	Sex            string
	Name           *PersonName
//	Occus          []string
	Occus          map[string]string // obliged to have a map because of go unmarshalling / php json_encode(JSON_FORCE_OBJECT)
	Birth          *Event
	Death          *Event
	Ids_in_sources interface{} // map[string]string
	Partial_ids    interface{} // map[string]string
	Trust          string
	Issues         []string
	Notes          []string
}

type PersonName struct {
	Usual     string
	Given     string
	Family    string
//	Spouse    []string
	Spouse    interface{}
	Official  *OfficialName
	Fame      *FameName
	Nicknames []string
//	Alter     []string
	Alter     interface{}
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

type PersonHistoryEntry struct {
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
var Partial_ids_labels = map[string]string{
	"lerrcp":   "Gauquelin",
	"muller":   "Müller",
	"cpara":    "Comité Para",
	"csicop":   "CSICOP",
	"cfepp":    "CFEPP",
	"ertel":    "Ertel",
	"g55":      "Gauquelin 1955",
	"wd":       "Wikidata",
}

// to decode view_issue_person
type PersonIssue struct {
    Slug        string  `json:"issue_slug"`
    Type        string  `json:"issue_type"`
    Description string  `json:"issue_description"`
}

// ************************** PersonName *******************************

func (p *Person) String() string {
	return p.Slug
}

/**
    Returns a string representing a person name
**/
func (n *PersonName) DisplayedName() string {
	if n.Fame.Full != "" {
		return n.Fame.Full
	}
	if n.Fame.Given != "" && n.Fame.Family != "" {
		return n.Fame.Given + " " + n.Fame.Family
	}
	//
	var family string
	if n.Fame.Family != "" {
	    family = n.Fame.Family
	} else {
        family = n.Family
        if n.Nobl != "" {
            if string(n.Nobl[len(n.Nobl)-1]) == "'" {
                family = n.Nobl + family
            } else {
                family = n.Nobl + " " + family
            }
        }
	}
	//
	var given string
	if n.Fame.Given != "" {
	    given = n.Fame.Given
	} else {
	    given = n.Given
	}
	//
	if given == "" {
	    // happens for Gauquelin persons with name not correctly restored
	    return family
	}
    return given + " " + family
}

// ************************** Get one *******************************

func GetPersonFromSlug(restURL string, slug string) (person *Person, err error) {
	url := restURL + "/person?slug=eq." + slug
	return GetPerson(url)
}

func GetPersonFromId(restURL string, id int) (person *Person, err error) {
	url := restURL + "/person?id=eq." + strconv.Itoa(id)
	return GetPerson(url)
}

func GetPerson(url string) (person *Person, err error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding person data "+url)
	}
	persons := []Person{}
	if err = json.Unmarshal(responseData, &persons); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal person data "+url+"\n"+string(responseData)+"\n")
	}
	if len(persons) > 1 {
		return nil, werr.New("One person expected - Several persons found at url"+url)
	}
	if len(persons) == 0 {
		// do not return error because should be handled as 404 by caller
		return nil, werr.New("Person not found at url "+url)
	}
	return &persons[0], nil
}

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
    /////////// TODO replace by slugify ///////////
    str = strings.ReplaceAll(str, " ", "-")
    str = strings.ToLower(str)
    /////////// end TODO replace by slugify ///////////
	url := restURL + "/search?slug=like." + str + "*&order=slug"
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

// ************************** Compute fields *******************************

/** Computes field 'Groups' of a person **/
func (p *Person) ComputeGroups(restURL string) (err error) {
	url := restURL + "/view_person_groop?person_id=eq." + strconv.Itoa(p.Id)
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


/** Computes field 'Isues' of a person **/
func (p *Person) ComputeIssues(restURL string) (err error) {
	url := restURL + "/view_issue?person_id=eq." + strconv.Itoa(p.Id)
	response, err := http.Get(url)
	if err != nil {
		return werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return werr.Wrapf(err, "Error decoding PersonGroup data"+string(responseData)+"\n")
	}
	if err = json.Unmarshal(responseData, &p.Issues); err != nil {
		return werr.Wrapf(err, fmt.Sprintf("Error json Unmarshal Issues\n%s\n", string(responseData)))
	}
	return nil
}

/** Returns a boolean indicating if a person has a birth certificate **/
func (p *Person) HasBC() bool {
	return p.Acts != nil
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

// ************************** Raw ordering *******************************
// Auxiliary code to sort the raw fields of a person

// Used by ctxt.template
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
	case "3a_sports":
		return RawPersonSortedFields["3a_sports"]
	case "csi":
		return RawPersonSortedFields["csi"]
	case "cfepp-final3":
		return RawPersonSortedFields["cfepp-final3"]
	default:
		return []string{}
	}
}

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
		"3a_sports": {
			"L",
			"MF",
			"NR",
			"ENG",
			"MA_",
			"G_NR",
			"MA12",
			"MARS",
			"NAME",
			"PUBL",
			"QUEL",
			"AUFAB",
			"CSINR",
			"KURTZ",
			"PHAS_",
			"STUND",
			"CFEPNR",
			"EXTEND",
			"NATION",
			"TAGMON",
			"ZITATE",
			"ZITSUM",
			"INDGRUP",
			"PARA_NR",
			"VORNAME",
			"ZITRANG",
			"GAUQ1955",
			"GEBDATUM",
			"GQBECORR",
			"NIENCORR",
			"NIENHUYS",
			"SPORTART",
			"CHRISNAME",
			"ZITSUM_OD",
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
		"cfepp-final3": {
			"S",
			"LT",
			"LV",
			"TR",
			"UT",
			"LAT",
			"SRC",
			"LONG",
			"NAME",
			"SPORT",
			"LOC_DATE",
			"UNIV_DATE",
			"BIRTH_PLACE",
			"POSTAL_CODE",
		},
	}
}
