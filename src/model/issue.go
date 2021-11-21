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
//	"sort"
)

type Issue struct {
	key            int
	Value          string
}
// ************************** Get many *******************************

/**
    Returns all issues, sorted by name.
**/
func GetIssues(restURL string) (issues []*Issue, err error) {
    
	url := restURL + "/api_issues"
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding issues data")
	}
	if err = json.Unmarshal(responseData, &issues); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal issues data")
	}
	return issues, nil
}

