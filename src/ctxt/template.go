/******************************************************************************
    Templates

    TODO - fix nl2br2 and sources.html {{url2href .Description | nl2br2 | safeHTML}}

    @license    GPL
    @history    2021-07-13 23:55:41+02:00, Thierry Graff : Creation
********************************************************************************/
package ctxt

import (
	"encoding/json"
	"html/template"
	"openg.local/openg/generic/tiglib"
	"openg.local/openg/model"
	"path/filepath"
	"time"
	"unicode"
)

// used to fill Context.Template
var tmpl *template.Template

func init() {
	var fmap = template.FuncMap{
		// Generic pipelines
		"dateIso":      dateIso,
		"modulo":       modulo,
		"safeHTML":     safeHTML,
		"ucFirst":      ucFirst,
		"numberFormat": numberFormat,
		"prettyPrint":  prettyPrint,
		// Pipelines related to current program
		"sourceNameFromSlug":    sourceNameFromSlug,
		"groupNameFromSlug":     groupNameFromSlug,
		"rawPersonSortedFields": rawPersonSortedFields,
	}
	tmpl = template.
		Must(template.
			New("").
			Funcs(fmap).
			ParseGlob(filepath.Join("view", "*.html"))).
		Option("missingkey=error")
	//	tmpl.New("listeActeurs").Funcs(fmap).ParseFiles(filepath.Join("view", "common", "listeActeurs.html"))
}

// ************************* Generic pipelines ********************************

func modulo(i, mod int) int {
	return i % mod
}

// Returns a date YYYY-MM-DD
func dateIso(t time.Time) template.HTML {
	return template.HTML(tiglib.DateIso(t))
}

// from https://www.php2golang.com/method/function.ucfirst.html
func ucFirst(str string) template.HTML {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return template.HTML(u + str[len(u):])
	}
	return template.HTML("")
}

func safeHTML(str string) template.HTML {
	return template.HTML(str)
}

func numberFormat(n int) template.HTML {
	return template.HTML(tiglib.NumberFormat(n, ' '))
}

// ************************* Pipelines related to current program ********************************

/**
    Returns the name of an information source from its slug
**/
func sourceNameFromSlug(slug string) template.HTML {
	name, _ := model.GetSourceNameFromSlug(config.RestURL, slug)
	return template.HTML(name)
}

/**
    Returns the name of a group from its slug
**/
func groupNameFromSlug(slug string) template.HTML {
	name, _ := model.GetGroupNameFromSlug(config.RestURL, slug)
	return template.HTML(name)
}

/**
    Returns the fields of a person for a given source
**/
func rawPersonSortedFields(sourceSlug string) []string {
	return model.GetRawPersonSortedFields(sourceSlug)
}

/**
    From https://siongui.github.io/2016/01/30/go-pretty-print-variable/
**/
func prettyPrint(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		return string(b)
	}
	return ""
}
