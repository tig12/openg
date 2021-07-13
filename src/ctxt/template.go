/******************************************************************************
    Templates

    @license    GPL
    @history    2021-07-13 23:55:41+01:00, Thierry Graff : Creation
********************************************************************************/
package ctxt

import (
	"openg.local/openg/generic/tiglib"
	"html/template"
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

// used to fill Context.Template
var tmpl *template.Template

func init() {
	var fmap = template.FuncMap{
		"dateIso":                      dateIso,
		"modulo":                       modulo,
		"nl2br":                        nl2br,
		"safeHTML":                     safeHTML,
		"ucFirst":                      ucFirst,
	}
	tmpl = template.
		Must(template.
			New("").
			Funcs(fmap).
			ParseGlob(filepath.Join("view", "*.html"))).
            Option("missingkey=error")
//	tmpl.New("listeActeurs").Funcs(fmap).ParseFiles(filepath.Join("view", "common", "listeActeurs.html"))
}

// ************************* pipelines ********************************

func modulo(i, mod int) int {
    return i % mod;
}

func nl2br(t string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(t), "\n", "<br>", -1))
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
