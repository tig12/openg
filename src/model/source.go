/******************************************************************************
    Information source

    @license    GPL
    @history    2021-07-17 17:29:31+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
)

type Source struct {
	Slug        string
	Type        string
	Name        string
	Title       string
	Authors     []string
	Edition     string
	Isbn        string
	Description template.HTML
	Parents     []string
}

// ************************** slug - names *******************************

// map source slug => source name
var sourceSlugNames = make(map[string]string)

func GetSourceSlugNames(restURL string) (map[string]string, error) {
	// lazy loading
	if len(sourceSlugNames) == 0 {
		sources, err := GetSources(restURL)
		if err != nil {
			return nil, werr.Wrapf(err, "Error calling GetSources()")
		}
		for _, source := range sources {
			sourceSlugNames[source.Slug] = source.Name
		}
	}
	return sourceSlugNames, nil
}

func GetSourceNameFromSlug(restURL, slug string) (string, error) {
	tmp, err := GetSourceSlugNames(restURL)
	if err != nil {
		return "", werr.Wrapf(err, "Error calling GetSourceSlugNames()")
	}
	return tmp[slug], nil
}

// ************************** Get many *******************************

func GetSources(restURL string) (sources []*Source, err error) {
	url := restURL + "/source"
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgres API: "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding sources data")
	}
	sources = []*Source{}
	if err = json.Unmarshal(responseData, &sources); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal sources data\n"+string(responseData)+"\n")
	}
	return sources, nil
}

// ************************** Get fields *******************************

func (s *Source) String() string {
	return s.Name
}
