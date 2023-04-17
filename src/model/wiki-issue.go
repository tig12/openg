/******************************************************************************
    Occupation

    @license    GPL
    @history    2021-07-17 17:29:31+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"openg.local/openg/generic/wilk/werr"
	"errors"
	"math"
	"strconv"
	"strings"
	"encoding/json"
	"io/ioutil"
	"net/http"
    "golang.org/x/text/cases"
    "golang.org/x/text/language"
)

type Issue struct {
	IssueSlug           string              `json:"issue_slug"`
	IssueType           string              `json:"issue_type"`
	IssueDescription    string              `json:"issue_description"`
	PersonSlug          string              `json:"person_slug"`
	PersonName          PersonName          `json:"person_name"`
	PersonBirth         Event               `json:"person_birth"`
	PersonIdsPartial    map[string]string   `json:"person_partial_ids"`
	PersonOccus         []string            `json:"person_occus"`
	WPName              string              `json:"wp_name"`
	WPSlug              string              `json:"wp_slug"`
}

func GetIssueTypeLabels() map[string]string{
    return map[string]string{
        "date": "Date or time problem",
        "name": "Problem in the name of the person",
        "tzo": "Timezone offset (Universal Time computation) problem",
    }
}

func GetNbIssues(restURL string) (int, error) {
	url := restURL + "/stats?select=n_issues"
	response, err := http.Get(url)
	if err != nil {
		return 0, werr.Wrapf(err, "Error calling postgrest API: \n"+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, werr.Wrapf(err, "Error decoding data containing nb issues")
	}
	type Info map[string]int
	tmp := []Info{}
	if err = json.Unmarshal(responseData, &tmp); err != nil {
		return 0, werr.Wrapf(err, "Error json Unmarshal nb issues\n"+string(responseData)+"\n")
	}
	if len(tmp) == 0 {
		return 0, werr.Wrapf(err, "EMPTY nb issues - need to be initialized")
	}
	return tmp[0]["n_issues"], nil
}

// ************************** Get many *******************************

/**
    Returns all issues for a given page, sorted by name.
    @param  page    Current page of the issues to load
    @param  limit   Nb of members to fetch
    @return issues
            nIssues Total number of issues
            pagemax Max nb of pages (used for navigation)
**/
func GetIssues(restURL string, page, limit int) (issues []*Issue, nIssues, pagemax int, err error) {
	nIssues, err = GetNbIssues(restURL)
	if err != nil {
		return nil, 0, 0, werr.Wrapf(err, "Cannot compute nIssues")
	}
	//
	pagemax = int(math.Ceil(float64(nIssues) / float64(limit)))
	if page > pagemax {
		page = pagemax
	}
	offset := (page - 1) * limit
	//
	url := restURL + "/view_issue" +
		"?limit=" + strconv.Itoa(limit) +
		"&offset=" + strconv.Itoa(offset)
	response, err := http.Get(url)
	if err != nil {
		return nil, 0, 0, werr.Wrapf(err, "Error calling postgrest API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, 0, werr.Wrapf(err, "Error decoding issues data")
	}
	if err = json.Unmarshal(responseData, &issues); err != nil {
		return nil, 0, 0, werr.Wrapf(err, "Error json Unmarshal issues data\n"+string(responseData)+"\n")
	}
	return issues, nIssues, pagemax, nil
}

// ************************** Slug manipulation *******************************

/** 
    Hack : compute the person name, slug and birth day from the issue slug
    Done to avoid a supplementary join on table person in the definition of view_wikiproject_issue
**/
func (i *Issue) ComputePersonFromSlug()(err error) {
    tmp1 := strings.Split(i.IssueSlug, "--")
    if len(tmp1) != 2 {
        return errors.New("Invalid issue slug: " + i.IssueSlug)
    }
    if len(tmp1[0]) < 11 {
        return errors.New("Invalid issue slug: " + i.IssueSlug)
    }
    i.PersonSlug = tmp1[0]
    // name
    personName := tmp1[0][:len(tmp1[0])-11]
    personName = strings.ReplaceAll(personName, "-", " ")
    personName = cases.Title(language.English, cases.Compact).String(personName)
    i.PersonName = PersonName{Usual:personName}
    // birth
    personDay := tmp1[0][len(tmp1[0])-10:]
    i.PersonBirth = Event{Date:personDay}
    return nil
}
