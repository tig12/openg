/******************************************************************************
    Person

    @license    GPL
    @history    2021-07-14 15:48:55+01:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
    "reflect"
    "net/http"
    "io/ioutil"
	"encoding/json"
	"openg.local/openg/generic/wilk/werr"
"fmt"
)


type Person struct {
    Id              int
    Slug            string
    To_check        bool
    Sources         []string
    Ids_in_sources  map[string]string
    Trust           string
    Trust_details   []string
    Sex             string
    Name            PersonName
    Occus           []string
    Birth           Event
    Death           Event
    Raw             map[string]map[string]string
    History         []HistoryEntry
    Notes           []string
}

type PersonName struct{
    Fame            string
    Given           string
    Usual           string
    Family          string
    Spouse          string
    Official        OfficialName
    Nicknames       []string
    Alternative     []string
    NobiliaryParticle   bool `json:"nobiliary-particle"`
}

type OfficialName struct{
    Given           string
    Family          string
}

type Event struct{
    Tzo             string
    Date            string
    Note            string
    Place           Place
    DateUT          string `json:"date-ut"`
}

type HistoryEntry struct{
	Details interface{} // TODO 
}

// ************************** Get one *******************************

func GetPerson(slug string) (p *Person, err error) {
    url := "http://localhost:1960/person?slug=eq." + slug
    response, err := http.Get(url)
    if err != nil {
		return nil, werr.Wrapf(err, "Error calling " + url)
    }    
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
		return nil, werr.Wrapf(err, "Error decoding person data " + slug)
    }
    persons := []Person{}
	if err = json.Unmarshal(responseData, &persons); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal person data " + slug)
	}
	if len(persons) > 1 {
		return nil, werr.Wrapf(err, "Several persons with identical slug: " + slug)
	}
	return &persons[0], nil
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
    if reflect.DeepEqual(p.Name, PersonName{}){
	    return "XXX" // TODO 
	}
	if p.Name.Usual != "" {
	    return p.Name.Usual
	}
	return p.Name.Given + " " + p.Name.Family
}


/** 
    Returns a person's birth day, format YYYY-MM-DD
    If Birth.Date exists, uses it.
    Otherwise uses field Birth.DateUT
**/
func (p *Person) GetBirthDay() string {
fmt.Printf("%+v\n",p.Birth.Date)
    if reflect.DeepEqual(p.Birth, Event{}){
//	    return "XXXX-XX-XX" // TODO 
	}
	var date string
	if p.Birth.Date != "" {
	    date = p.Birth.Date
	} else if p.Birth.DateUT != "" {
	    date = p.Birth.DateUT
	} else {
	    return "XXXX-XX-XX"
	}
	return date[:10]
}


