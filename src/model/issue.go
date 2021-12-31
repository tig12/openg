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
	PersonSlug string     `json:"slug"`
	PersonName PersonName `json:"name"`
	Values     []string   `json:"issues"`
}

// ************************** Get many *******************************

/**
    Returns all issues, sorted by name.
**/
func GetIssues(restURL string) (issues []*Issue, err error) {

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
