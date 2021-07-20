/******************************************************************************
    Information source

    @license    GPL
    @history    2021-07-17 17:29:31+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
    "net/http"
    "io/ioutil"
	"encoding/json"
	"openg.local/openg/generic/wilk/werr"
//"fmt"
)

type Source struct{
    Slug        string
    Type        string
    Name        string
    Title       string
    Authors     []string
    Edition     string
    Isbn        string
    Description string
    Parents     []string
}

// ************************** Get many *******************************

func GetSources() (sources []*Source, err error) {
    url := "http://localhost:1960/source"
    response, err := http.Get(url)
    if err != nil {
		return nil, werr.Wrapf(err, "Error calling " + url)
    }    
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
		return nil, werr.Wrapf(err, "Error decoding sources data")
    }
    sources = []*Source{}
	if err = json.Unmarshal(responseData, &sources); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal sources data")
	}
	
	// return the sources in an arbitrary order
	order := []string{
        "g5",
        "cura5",
        "newalch",
        "a1-booklet",
        "a1",
        "a2-booklet",
        "a2",
        "a3-booklet",
        "a3",
        "a4-booklet",
        "a4",
        "a5-booklet",
        "a5",
        "a6-booklet",
        "a6",
        "d6-booklet",
        "d6",
        "d10-booklet",
        "d10",
        "e1-booklet",
        "e1",
        "e3-booklet",
        "e3",
        "afd",
        "afd1-booklet",
        "afd1",
        "afd1-100",
        "afd5-booklet",
        "afd5",
        "csicop-committee",
        "csi",
        "si42",
	}
	var res = []*Source{}
	for _, slug := range(order){
	    for _, s := range(sources){
	        if s.Slug == slug {
	            res = append(res, s)
	            break
	        }
	    }
	}
	return res, nil
}

// ************************** Get fields *******************************

func (s *Source) String() string {
    return s.Name
}

    