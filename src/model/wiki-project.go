/******************************************************************************
    Person

    @license    GPL
    @history    2021-07-14 15:48:55+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"openg.local/openg/generic/wilk/werr"
	"sort"
    "strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type WikiProject struct {
	Header      DocumentHeader `json:"header"`
	Id          int
	Slug        string
	Name        string
	Description string
	// Not in database
	Persons     []*WikiProjectPerson
	Issues      []*Issue  
}

/** 
    Partial data of a person
**/
type WikiProjectPerson struct {
    Id      int     `json:"person_id"`
    Slug    string  `json:"person_slug"`
	Name    *PersonName `json:"person_name"`
	Birth   *Event `json:"person_birth"`
}

// ************************** Get one *******************************

/**
    Loads a wiki project from database.
    Loads only the fields present in table wikiproject.
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
    Computes a wiki project with its fields Issues and Persons.
**/
func GetWikiProjectFullFromSlug(restURL, slug string) (project *WikiProject, err error) {
    project, err = GetWikiProjectFromSlug(restURL, slug)
	if err != nil {
		return project, werr.Wrapf(err, "Error calling GetWikiProjectFromSlug()")
	}
    err = project.ComputeActs(restURL)
	if err != nil {
		return project, werr.Wrapf(err, "Error calling project.ComputeActs()")
	}
    err = project.ComputeIssues(restURL)
	if err != nil {
		return project, werr.Wrapf(err, "Error calling project.ComputeIssues()")
	}
    return project, nil
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

/**
    Returns the wiki projects related to a birth certificate.
**/
func GetWikiProjectsOfBC(restURL string, bc *BC) (result []*WikiProject, err error) {
    if bc.OpenGauquelin == nil {
        return result, nil // empty
    }
    if bc.OpenGauquelin.WikiProjects == nil {
        return result, nil // empty
    }
    for _, wpSlug := range(*bc.OpenGauquelin.WikiProjects){
        wikiproject, err := GetWikiProjectFromSlug(restURL, wpSlug)
        if err != nil {
            return result, werr.Wrapf(err, "Error calling GetWikiProjectFromSlug("+wpSlug+")\n") // empty
        }
        result = append(result, wikiproject)
    }
	return result, nil
}

// ************************** Instance methods *******************************

/**
    Computes the persons related to a wiki project.
**/
func (wp *WikiProject) ComputeActs(restURL string) error {
	url := restURL + "/view_wikiproject_act?project_id=eq." + strconv.Itoa(wp.Id)
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

/**
    Computes the issues related to a wiki project.
**/
func (wp *WikiProject) ComputeIssues(restURL string) (err error) {
	url := restURL + "/view_wikiproject_issue?project_id=eq." + strconv.Itoa(wp.Id)
	response, err := http.Get(url)
	if err != nil {
		return werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return werr.Wrapf(err, "Error decoding wiki project persons data "+wp.Slug)
	}
	if err = json.Unmarshal(responseData, &wp.Issues); err != nil {
		return werr.Wrapf(err, "Error json Unmarshal wiki project issues data "+wp.Slug+"\n"+string(responseData)+"\n")
	}
	for _, issue := range(wp.Issues) {
	    err = issue.ComputePersonFromSlug() // hack see comment of ComputePersonFromSlug()
        if err != nil {
            return werr.Wrapf(err, "Error calling issue.ComputePersonFromSlug() ")
        }
	}
	sortedIssues := make(issueSlice, 0, len(wp.Issues))
	for _, elt := range wp.Issues {
		sortedIssues = append(sortedIssues, elt)
	}
    sort.Sort(sortedIssues)
    wp.Issues = sortedIssues
    return nil
}    

// Auxiliaries of ComputeIssues() to sort issues by person slug
type issueSlice []*Issue
func (issues issueSlice) Len() int           { return len(issues) }
func (issues issueSlice) Less(i, j int) bool { return issues[i].PersonSlug < issues[j].PersonSlug }
func (issues issueSlice) Swap(i, j int)      { issues[i], issues[j] = issues[j], issues[i] }
