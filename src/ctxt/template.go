/******************************************************************************
    Templates

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
	"strings"
	"time"
	"unicode"
)

// used to fill Context.Template
var tmpl *template.Template

func init() {
	var fmap = template.FuncMap{
		// Generic pipelines
		"add":             add,
		"dateIso":         dateIso,
		"modulo":          modulo,
		"nl2br":           nl2br,
		"numberFormat":    numberFormat,
		"prettyPrint":     prettyPrint,
		"safeHTML":        safeHTML,
		"ucFirst":         ucFirst,
		"whiteSpace2nbsp": whiteSpace2nbsp,
		// Pipelines related to current program
		"countryLabel":          countryLabel,
		"groupNameFromSlug":     groupNameFromSlug,
		"sourceNameFromSlug":    sourceNameFromSlug,
		"rawPersonSortedFields": rawPersonSortedFields,
		"trustDescription":      trustDescription,
	}
	tmpl = template.
		Must(template.
			New("").
			Funcs(fmap).
			ParseGlob(filepath.Join("view", "*.html"))).
		Option("missingkey=error")
//		template.Must(template.ParseGlob(filepath.Join("view", "wiki", "*.html")))
}

// ************************* Generic pipelines ********************************

func add(i, j int) int {
	return i + j
}

func modulo(i, mod int) int {
	return i % mod
}

// Returns a date YYYY-MM-DD
func dateIso(t time.Time) template.HTML {
	return template.HTML(tiglib.DateIso(t))
}

func nl2br(t string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(t), "\n", "<br>", -1))
}

func numberFormat(x interface{}) template.HTML {
	return template.HTML(tiglib.NumberFormat(x, ' '))
}

/**
    From https://siongui.github.io/2016/01/30/go-pretty-print-variable/
**/
func prettyPrint(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
	    str := string(b)
	    // dirty hack to have a href link in person's history raw value (person-build.html)
	    str = strings.Replace(str, "\\u003c", "<", -1)
	    str = strings.Replace(str, "\\u003e", ">", -1)
	    str = strings.Replace(str, "\\\"", "\"", -1)
		return str
	}
	return ""
}

func safeHTML(str string) template.HTML {
	return template.HTML(str)
}

// from https://www.php2golang.com/method/function.ucfirst.html
func ucFirst(str string) template.HTML {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return template.HTML(u + str[len(u):])
	}
	return template.HTML("")
}

func whiteSpace2nbsp(t string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(t), " ", "&nbsp;", -1))
}

// ************************* Pipelines related to current program ********************************

// Returns the name of an information source from its slug
func countryLabel(countryCode string) template.HTML {
    var name string
    if _, ok := model.CountryCodesNames[countryCode]; ok {
        name = model.CountryCodesNames[countryCode]
    } else {
        name = "???"
    }
	return template.HTML(name)
}

// Returns the name of an information source from its slug
func sourceNameFromSlug(slug string) template.HTML {
	name, _ := model.GetSourceNameFromSlug(config.RestURL, slug)
	return template.HTML(name)
}

// Returns the name of a group from its slug
func groupNameFromSlug(slug string) template.HTML {
	name, _ := model.GetGroupNameFromSlug(config.RestURL, slug)
	return template.HTML(name)
}

// Returns the fields of a person for a given source
func rawPersonSortedFields(sourceSlug string) []string {
	return model.GetRawPersonSortedFields(sourceSlug)
}

// Returns the meaning of a trust level 
// @param  str  The trus level, a string between "1" to "5".
func trustDescription(str string) template.HTML {
    var res string
	switch str {
	case "1":
		res = "1 = Related to a Hospital Certificate - reliable"
    break
	case "2":
		res = "2 = Related to a Birth Certificate - reliable"
    break
	case "3":
		res = "3 = Related to a Birth Record - may contain mistakes"
    break
	case "4":
		res = "4 = Related to an official document ; birth time missing or needing to be checked"
    break
	case "5":
		res = "5 = To check - not related to an official document"
    break
    }
	return template.HTML(res)
}
