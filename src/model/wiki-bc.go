/******************************************************************************
    BC = Birth certificate

    @license    GPL
    @history    2022-02-16 01:51:37+01:00, Thierry Graff : Creation
********************************************************************************/
package model

import ()

type BC struct {
	Header DocumentHeader `json:"header"`
	Source struct {
		CivilRegistry struct {
			Name    string `json:"name"`
			Place   string `json:"place"`
			Country string `json:"cy"`
			C1      string `json:"c1"`
			C2      string `json:"c2"`
			C3      string `json:"c3"`
			Web     struct {
				URL  string `json:"url"`
				Page string `json:"page"`
			} `json:"web"`
		} `json:"civil-registry"`
		DocumentCreation struct {
			Date  string `json:"date"`
			Place string `json:"place"`
		} `json:"document-creation"`
		Notes  string            `json:"notes"`
		Images map[string]string `json:"images"`
	}
	OpenGauquelin struct {
		Projects map[string]string `json:"projects"`
	} `json:"opengauquelin"`
	Person PartialPerson `json:"person"`
	Extras PartialPerson `json:"extras"`
}

/**
    Returns the urls of the images representing the BC.
**/
func ComputeBCImageURLs(p *Person, baseUrl string) (result []string) {
	if !p.HasBC() {
		return result
	}
	//
	candidates := map[string]string{}
	if len(p.Acts.Birth.Source.Images) == 0 {
		candidates["0"] = "BC.jpg"
	} else {
		candidates = p.Acts.Birth.Source.Images
	}
	for _, candidate := range candidates {
		// "wiki-data" is hard coded, see run-openg.go
		url := "/wiki-data/birth/" + Slug2URL(p) + "/" + candidate
		result = append(result, url)
	}
	return result
}
