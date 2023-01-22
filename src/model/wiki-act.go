/******************************************************************************
    Act = official document like birth, mariage and death certificates.

    @license    GPL
    @history    2022-02-16 01:51:37+01:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"strings"
)

/**
    Struct used to store field Acts of a person
**/
type Acts struct {
	Birth *BC `json:"birth"`
}

/**
    Returns the relative URL of a person directory within wiki-data.
**/
func Slug2URL(p *Person) (result string) {
	parts := strings.Split(p.Slug, "-")
	l := len(parts)
	// ex: file = 1897/11/26/accard-robert-1897-11-26/BC.yml
	result = strings.Join([]string{
		parts[l-3],
		parts[l-2],
		parts[l-1],
		p.Slug,
	}, "/")
	return result
}
