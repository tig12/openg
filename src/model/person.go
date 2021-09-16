/******************************************************************************
    Person

    @license    GPL
    @history    2021-07-14 15:48:55+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
//"fmt"
)

// see init function at the end

type Person struct {
	Id       int
	Slug     string
	Sex            string
	Name           PersonName
	Occus          []string
	Birth          Event
	Death          Event
	Sources        []string
	//	Sources        []*Source
	Ids_in_sources map[string]string
	Trust          interface{}
	Acts           []string
	History        []HistoryEntry
	Todo           []string
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
	Raw     map[string]string
	New     interface{}
}

// ************************** Get one *******************************

/**
    Loads a person from database
**/
func GetPerson(restURL, slug string) (person *Person, err error) {
	url := restURL + "/person?slug=eq." + slug
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding person data "+slug)
	}
	persons := []Person{}
	if err = json.Unmarshal(responseData, &persons); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal person data "+slug)
	}
	if len(persons) > 1 {
		return nil, werr.New("Several persons with identical slug: "+slug)
	}
	if len(persons) == 0 {
		return nil, werr.New("Unexisting person with slug: "+slug)
	}
	return &persons[0], nil
}

// ************************** Get many *******************************

func GetPersons(restURL string) (p []*Person, err error) {
	url := restURL + "/person?limit=10&offset=0"
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding persons data")
	}
	persons := []*Person{}
	if err = json.Unmarshal(responseData, &persons); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal persons data")
	}
	return persons, nil
}

// ************************** Get fields *******************************

func (p *Person) String() string {
	return p.Slug
}

/**
    Returns field UsualName if it exists.
    Otherwise, returns a concatenation of given and family name.
**/
func (p *Person) GetName() string {
	if p.Name.Usual != "" {
		return p.Name.Usual
	}
	if p.Name.Given == "" {
		return p.Name.Family
	}
	return p.Name.Given + " " + p.Name.Family

}

/**
    Returns a person's birth date (day and time), format YYYY-MM-DD HH:MM:SS
    If Birth.Date exists, uses it.
    Otherwise uses field Birth.DateUT
**/
func (p *Person) GetBirthDate() string {
	return GetBirthDate(p.Birth.Date, p.Birth.DateUT)
}

/**
    Function used by Person and GroupMember
**/
func GetBirthDate(date, dateUT string) (res string) {
	if date != "" {
		res = date
	} else if dateUT != "" {
		res = dateUT
	} else {
		return "XXXX-XX-XX XX:XX:XX"
	}
	return res
}

/**
    Returns a person's birth day, format YYYY-MM-DD
    If Birth.Date exists, uses it.
    Otherwise uses field Birth.DateUT
**/
func (p *Person) GetBirthDay() string {
	return p.GetBirthDate()[:10]
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
