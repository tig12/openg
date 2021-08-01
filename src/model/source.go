/******************************************************************************
    Information source

    @license    GPL
    @history    2021-07-17 17:29:31+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
	"html/template"
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
		return nil, werr.Wrapf(err, "Error calling "+url)
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
	    // primary
	    "lerrcp",
	    "afd",
		"csicop-committee",
		// secondary
		"cura5",
		"newalch",
		"g5",
		// Gauquelin
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
		// Müller
		"afd1-booklet",
		"afd1",
		"afd1-100",
		"afd5-booklet",
		"afd5",
		// csicop
		"csi",
		"si42",
	}
	var res = []*Source{}
	for _, slug := range order {
		for _, s := range sources {
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
