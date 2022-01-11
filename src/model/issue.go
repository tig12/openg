/******************************************************************************
    Occupation

    @license    GPL
    @history    2021-07-17 17:29:31+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
)

type Issue struct {
	PersonSlug  string            `json:"slug"`
	PersonName  PersonName        `json:"name"`
	PersonBirth Event             `json:"birth"`
	IdsPartial  map[string]string `json:"ids_partial"`
	PersonOccus []string          `json:"occus"`
	Values      []string          `json:"issues"`
}

// ************************** Get many *******************************

/**
    Returns all issues for a given page, sorted by name.
    @param  page    Current page of the issues to load
    @param  limit   Nb of members to fetch
**/
func GetIssues(restURL string, page, limit int) (issues []*Issue, err error) {

/* 
    n:= GetNbIssues(restURL)
	pagemax := int(math.Ceil(float64(n) / float64(limit)))
	group.NPages = pagemax
	if page > pagemax {
		page = pagemax
	}
	offset := (page - 1) * limit
*/
	
	url := restURL + "/api_issue"
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding issues data")
	}
	if err = json.Unmarshal(responseData, &issues); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal issues data\n"+string(responseData)+"\n")
	}
	return issues, nil
}

func GetNbIssues(restURL string) int {
    return 261 // TODO 
}
