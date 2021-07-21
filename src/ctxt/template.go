/******************************************************************************
    Templates

    TODO - fix nl2br2 and sources.html {{url2href .Description | nl2br2 | safeHTML}}

    @license    GPL
    @history    2021-07-13 23:55:41+02:00, Thierry Graff : Creation
********************************************************************************/
package ctxt

import (
	"html/template"
	"openg.local/openg/generic/tiglib"
	"openg.local/openg/model"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// used to fill Context.Template
var tmpl *template.Template

func init() {
	var fmap = template.FuncMap{
		// Generic pipelines
		"dateIso": dateIso,
		"modulo":  modulo,
		"nl2br":   nl2br,
		// TODO hack to remove
		"nl2br2":   nl2br2,
		"safeHTML": safeHTML,
		"ucFirst":  ucFirst,
		"url2href": url2href,
		// Pipelines related to current program
		"sourceNameFromSlug":    sourceNameFromSlug,
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

func nl2br(t string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(t), "\n", "<br>", -1))
}

// TODO hack to remove
func nl2br2(t string) string {
	return strings.Replace(t, "\n", "<br>", -1)
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

// Adds a href links around words starting by http
// TODO fix to return a template.HTML
func url2href(str string) template.HTML {
	//func url2href(str string) string {
	var re = regexp.MustCompile(`\b(https?\:\/\/.*?)\s`)
	res := re.ReplaceAllString(str, `<a href="$1">$1</a>`)
	return template.HTML(res)
	//	return res
}

// ************************* Pipelines related to current program ********************************

/**
    Returns the name of an information source from its slug
    TODO find a cleaner way to grab config ?
**/
func sourceNameFromSlug(slug string) template.HTML {
	name, _ := model.GetSourceNameFromSlug(config.RestURL, slug)
	return template.HTML(name)
}

/**
    Returns the fields of a person for a given source
**/
func rawPersonSortedFields(sourceSlug string) []string {
	return model.GetRawPersonSortedFields(sourceSlug)
}
