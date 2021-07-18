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
	return sources, nil
}

// ************************** Get fields *******************************

func (s *Source) String() string {
    return s.Name
}

    