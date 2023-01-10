/******************************************************************************
    Informations about acts (birth, marriage, death certificates).

    @license    GPL
    @history    2022-02-16 01:51:37+01:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type BC struct { // implements Act
	Header   interface{}
	Source interface{}
	Person   Person
	OpenGauquelin interface{}
	Extras interface{}
}

/**
    Adds birth certificate in p.ActObjects["birth"]
    if p.Acts contains the string "birth"
**/
func ComputeBC(p *Person, dirActs string) error {
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
	return nil
}
