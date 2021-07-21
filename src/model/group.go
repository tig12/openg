/******************************************************************************
    Information source

    @license    GPL
    @history    2021-07-17 17:29:31+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"encoding/json"
	"io/ioutil"
//"io"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
//"fmt"
)

type Group struct {
    Id          int
	Slug        string
	Name        string
	Description string
	SourceSlugs     []string `json:"sources"`
	ParentSlugs     []string `json:"parents"`
	MemberSlugs     []string `json:"members"`
	// not stored in database
	Sources     []*Source
	Parents     []*Group
	Members     []*Person
}

// ************************** Get one *******************************

func GetGroupBySlug(restURL, slug string) (group *Group, err error) {
    var url string
    var responseData []byte
    var response *http.Response
    
    // get the group
	url = restURL + "/groop?slug=eq." + slug
	response, err = http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling "+url)
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding Group data")
	}
	var tmp []*Group
	if err = json.Unmarshal(responseData, &tmp); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal Group")
	}
	group = tmp[0]
	
    // get the persons of the group
	url = restURL + "/api_persongroop?group_slug=eq." + slug
	response, err = http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling "+url)
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding group members ("+slug+")")
	}
	if err = json.Unmarshal(responseData, &group.Members); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal Group.Members")
	}
	
	return group, nil
}

// ************************** Get fields *******************************

func (g *Group) String() string {
	return g.Name
}
