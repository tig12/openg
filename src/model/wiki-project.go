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
	//	"fmt"
)

type WikiProject struct {
	Header      DocumentHeader `json:"header"`
	Id          int
	Slug        string
	Name        string
	Description string
}

// ************************** Get one *******************************

/**
    Loads a wiki project from database
**/
func GetWikiProject(restURL, slug string) (project *WikiProject, err error) {
	url := restURL + "/api_wikiproject?slug=eq." + slug
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
