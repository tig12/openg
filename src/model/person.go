/******************************************************************************
    Person

    @license    GPL
    @history    2021-07-14 15:48:55+01:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"openg.local/openg/generic/wilk/werr"
    "net/http"
    "io/ioutil"
	"encoding/json"
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
    NobiliaryParticle   bool
    
}

type OfficialName struct{
    Given           string
    Family          string
}

type Event struct{
    tzo             string
    date            string
    note            string
    place           Place
    dateUt          string
}

type Place struct{
    c2              string
    c3              string
    cy              string
    lg              float64
    lat             float64
    name            string
    geoid           string
}

type HistoryEntry struct{
	Details interface{}
}
// ************************** Nom *******************************

func (p *Person) String() string {
	return p.Slug
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
fmt.Printf("%+v\n", p)
		return nil, werr.Wrapf(err, "Error json Unmarshal person data " + slug)
	}
	return p, nil
}

