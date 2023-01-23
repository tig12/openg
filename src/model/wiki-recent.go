/******************************************************************************
    A WikiRecent represents a recent addition to the database through the wiki

    @license    GPL
    @history    2023-01-23 18:43:20+01:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"openg.local/openg/generic/wilk/werr"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"
//"fmt"
)

type WikiRecent struct {
	IdPerson    int `json:"id_person"`
	Date        string `json:"dateadd"`
	Description string
	// Not in database
	Person     *Person
}

/** 
    Returns WikiRecent only with data stored in database
**/
func GetWikiRecents(restURL string, limit int) (recents []*WikiRecent, err error) {
	url := restURL + "/wikirecent?order=dateadd.desc&limit=" + strconv.Itoa(limit)
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding wikirecent data")
	}
	if err = json.Unmarshal(responseData, &recents); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal wikirecent data:\n"+string(responseData)+"\n")
	}
	return recents, nil
}

/** 
    Returns WikiRecent with full data
**/
func GetWikiRecentsFull(restURL string, limit int) ([]*WikiRecent, error) {
    recents, err := GetWikiRecents(restURL, limit)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling GetWikiRecents()")
	}
    for i, _ := range(recents) {
        err = recents[i].ComputePerson(restURL, recents[i].IdPerson)
        if err != nil {
            return nil, werr.Wrapf(err, "Error calling ComputePerson()")
        }
    }
    return recents, nil
}

// ************************** Compute fields *******************************

func (recent *WikiRecent) ComputePerson(restURL string, personId int) (err error) {
    recent.Person, err = GetPersonFromId(restURL, recent.IdPerson)
    if err != nil {
        return werr.Wrapf(err, "Error calling GetPersonFromId()")
    }
    return nil
}
