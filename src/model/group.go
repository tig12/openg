/******************************************************************************
    Group of persons

    @license    GPL
    @history    2021-07-17 17:29:31+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
)

const GROUP_TYPE_OCCU = "occu"
const GROUP_TYPE_HISTORICAL = "history"

type Group struct {
	Id            int
	Slug          string
	Name          string
	WD            string
	N             int
	Type          string
	Description   string
	Download      string
	SourceSlugs   []string `json:"sources"`
	ParentSlugs   []string `json:"parents"`
	ChildrenSlugs []string `json:"children"`
	MemberSlugs   []string `json:"members"`
	// not stored in database
	Sources  []*Source
	Parents  []*Group
	Children []*Group
	Members  []*GroupMember
}

/** Simplified representation of a person **/
type GroupMember struct {
	Slug           string `json:"person_slug"`
	Ids_in_sources map[string]string
	Ids_partial    map[string]string
	Sex            string
	Name           PersonName
	Occus          []string
	Birth          Event
}

// ************************** slug - names *******************************
// map group slug => group name
var groupSlugNames = make(map[string]string)

func GetGroupSlugNames(restURL string) (map[string]string, error) {
	// lazy loading
	if len(groupSlugNames) == 0 {
		groups, err := GetGroups(restURL)
		if err != nil {
			return nil, werr.Wrapf(err, "Error calling GetGroup()")
		}
		for _, group := range groups {
			groupSlugNames[group.Slug] = group.Name
		}
	}
	return groupSlugNames, nil
}

func GetGroupNameFromSlug(restURL, slug string) (string, error) {
	tmp, err := GetGroupSlugNames(restURL)
	if err != nil {
		return "", werr.Wrapf(err, "Error calling GetGroupSlugNames()")
	}
	return tmp[slug], nil
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
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding Group data")
	}
	var tmp []*Group
	if err = json.Unmarshal(responseData, &tmp); err != nil {
		return nil, werr.Wrapf(err, fmt.Sprintf("Error json Unmarshal Group '%s'\n%s\n", slug, string(responseData)))
	}
	if len(tmp) == 0 {
		var group = Group{}
		return &group, werr.Wrapf(err, "EMPTY Stats - need to be initialized")
	}
	group = tmp[0]

	// get the members of the group
	url = restURL + "/api_persongroop?group_slug=eq." + slug
	response, err = http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding group members ("+slug+")")
	}
	if err = json.Unmarshal(responseData, &group.Members); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal Group.Members\n" +string(responseData) + "\n")
	}

	return group, nil
}

// ************************** Get many *******************************

/**
    Returns all the groups present in database.
    Each group only contains the fields stored in database (no computed fields like members).
**/
func GetGroups(restURL string) (groups []*Group, err error) {
	url := restURL + "/groop"
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding groups data")
	}
	groups = []*Group{}
	if err = json.Unmarshal(responseData, &groups); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal groups data\n" +string(responseData) + "\n")
	}
	return groups, nil
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
