/******************************************************************************
    BC = Birth certificate

    @license    GPL
    @history    2022-02-16 01:51:37+01:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
// "gopkg.in/yaml.v3"
// "os"
// "strings"
//"fmt"
)

type BC struct {
	Header        DocumentHeader `json:"header"`
	Source        struct {
        CivilRegistry    interface{} `json:"civil-registry"`
        DocumentCreation struct{
            Date string `json:"date"`
            Place string `json:"place"`
        } `json:"document-creation"`
        Notes            string `json:"notes"`
        Images           map[string]string `json:"images"`
    }
	OpenGauquelin struct {
	    Projects    map[string]string `json:"projects"`
	} `json:"opengauquelin"`
	Person        PartialPerson `json:"person"`
	Extras        PartialPerson `json:"extras"`
}


/**
    Returns the urls of the images representing the BC.
**/
func ComputeBCImages(p *Person, baseUrl string) (result []string, err error) {
/* 
    candidates := []string{}
    if len(p.BC.Source.Images) == 0 {
        candidates = append(candidates, "BC.jpg")
    } else {
        candidates = p.BC.Source.Images
    }
fmt.Printf("p.BC.Source = %+v\n",p.BC.Source)
fmt.Printf("candidates = %+v\n",candidates)
*/
    
	/*
		found := false
		for _, act := range p.Acts {
			if act == "birth" {
				found = true
				break
			}
		}
		if !found {
			return nil
		}
		parts := strings.Split(p.Slug, "-")
		l := len(parts)
		// ex: file = /path/to/acts/birth/1897/11/26/accard-robert-1897-11-26/BC.yml
		file := strings.Join([]string{
			dirActs,
			"birth",
			parts[l-3],
			parts[l-2],
			parts[l-1],
			p.Slug,
			"BC.yml",
		}, string(os.PathSeparator))

		contents, err := os.ReadFile(file)
		if err != nil {
			if os.IsNotExist(err) {
				return nil // p.Acts remains empty
			}
			return err
		}
		var bc BC
		err = yaml.Unmarshal(contents, &bc)
		if err != nil {
			return err
		}
		p.Acts["birth"] = bc
	*/
	return result, nil
}
