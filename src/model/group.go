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
	Members     []*GroupMember
}

/** Simplified representation of a person **/
type GroupMember struct {
	Slug           string `json:"person_slug"`
	Ids_in_sources map[string]string
	Sex            string
	Name           PersonName
	Occus          []string
	Birth          Event
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
    if len(tmp) == 0 {
        var group = Group{}
		return &group, werr.Wrapf(err, "EMPTY Stats - need to be initialized")
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

/**
    Returns a person's birth date (day and time), format YYYY-MM-DD HH:MM:SS
    If Birth.Date exists, uses it.
    Otherwise uses field Birth.DateUT
**/
func (p *GroupMember) GetBirthDate() string {
	return GetBirthDate(p.Birth.Date, p.Birth.DateUT)
}