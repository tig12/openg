/******************************************************************************
    Person

    @license    GPL
    @history    2021-07-14 15:48:55+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
    "strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
//	"fmt"
)

type WikiProject struct {
	Header      DocumentHeader `json:"header"`
	Id          int
	Slug        string
	Name        string
	Description string
	// Not in database
	Persons     []*WikiProjectPerson
}

/** 
    Partial data of a person
**/
type WikiProjectPerson struct {
    Id      int     `json:"person_id"`
    Slug    string  `json:"person_slug"`
	Name    *PersonName `json:"name"`
	Birth   *Event `json:"birth"`
}

// ************************** Get one *******************************

/**
    Loads a wiki project from database
**/
func GetWikiProjectFromSlug(restURL, slug string) (project *WikiProject, err error) {
	url := restURL + "/wikiproject?slug=eq." + slug
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding wiki project data "+slug)
	}
	projects := []*WikiProject{}
	if err = json.Unmarshal(responseData, &projects); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal wiki project data "+slug+"\n"+string(responseData)+"\n")
	}
	if len(projects) == 0 {
		// do not return error because should be handled as 404 by caller
		return nil, nil
	}
	return projects[0], nil
}

/**
    Computes the wiki projects related to a birth certificate.
**/
func ComputeBCWikiProjects(restURL string, bc *BC) (result []*WikiProject, err error) {
    if bc.OpenGauquelin == nil {
        return result, nil // empty
    }
    if bc.OpenGauquelin.WikiProjects == nil {
        return result, nil // empty
    }
    for _, wpSlug := range(*bc.OpenGauquelin.WikiProjects){
        wikiproject, err := GetWikiProjectFromSlug(restURL, wpSlug)
        if err != nil {
            return []*WikiProject{}, werr.Wrapf(err, "Error calling GetWikiProjectFromSlug("+wpSlug+")\n")
        }
        result = append(result, wikiproject)
    }
	return result, nil
}

// ************************** Get many *******************************

func GetActiveWikiProjects(restURL string) (result []*WikiProject, err error) {
	url := restURL + "/wikiproject?status=eq.active"
	response, err := http.Get(url)
	if err != nil {
		return result, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, werr.Wrapf(err, "Error decoding active wiki projects")
	}
	if err = json.Unmarshal(responseData, &result); err != nil {
		return result, werr.Wrapf(err, "Error json Unmarshal wiki projects\n"+string(responseData)+"\n")
	}
    return result, nil
}

// ************************** Instance methods *******************************

/**
    Computes the persons related to a wiki project.
**/
func (wp *WikiProject) ComputePersons(restURL string) error {
	url := restURL + "/view_wikiproject_person?project_id=eq." + strconv.Itoa(wp.Id)
	response, err := http.Get(url)
	if err != nil {
		return werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return werr.Wrapf(err, "Error decoding wiki project persons data "+wp.Slug)
	}
	if err = json.Unmarshal(responseData, &wp.Persons); err != nil {
		return werr.Wrapf(err, "Error json Unmarshal wiki project persons data "+wp.Slug+"\n"+string(responseData)+"\n")
	}
    return nil
}    

