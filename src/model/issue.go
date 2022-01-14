/******************************************************************************
    Occupation

    @license    GPL
    @history    2021-07-17 17:29:31+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
	"strconv"
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
    @return issues
            nIssues Total number of issues
            pagemax Max nb of pages (used for navigation)
**/
func GetIssues(restURL string, page, limit int) (issues []*Issue, nIssues, pagemax int, err error) {

	nIssues, err = GetNbIssues(restURL)
	if err != nil {
		return nil, 0, 0, werr.Wrapf(err, "Cannot compute nIssues")
	}

	pagemax = int(math.Ceil(float64(nIssues) / float64(limit)))
	if page > pagemax {
		page = pagemax
	}
	offset := (page - 1) * limit

	url := restURL + "/api_issue" +
		"?limit=" + strconv.Itoa(limit) +
		"&offset=" + strconv.Itoa(offset)
	response, err := http.Get(url)
	if err != nil {
		return nil, 0, 0, werr.Wrapf(err, "Error calling postgres API: "+url)
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
