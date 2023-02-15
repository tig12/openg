/**

**/
package control

import (
	"fmt"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	"strings"
)

// To display the download form
type detailsDownloadForm struct {
	Occupations []*model.Group
}

func ShowDownloads2(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case "POST":
		//
		// Process form
		//
		data, err := processDownloadForm(r)
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", data)
		csv, err := model.ExportGroup(data)
		fmt.Printf("%+v\n", csv)
		if err != nil {
			return err
		}
		//
		redirect := "/downloads2"
		ctx.Redirect = redirect
		return nil
	default:
		//
		// Show form
		//
		occupations, err := model.GetOccus(ctx.Config.RestURL)
		if err != nil {
			return err
		}

		ctx.TemplateName = "download2.html"
		ctx.Page = &ctxt.Page{
			Header: ctxt.Header{
				Title:    "Download | Open Gauquelin database",
				CSSFiles: []string{"/static/css/form.css"},
				//JSFiles: []string{},
			},
			Details: detailsDownloadForm{
				Occupations: occupations,
			},
		}
		return nil
	}
}

func processDownloadForm(r *http.Request) (*model.DownloadFormData, error) {
	data := &model.DownloadFormData{} // see model/export.go
	var err error
	if err = r.ParseForm(); err != nil {
		return data, err
	}
	//
	data.What = r.PostFormValue("radio-what")
	if strings.HasPrefix(data.What, "occu-") {
		data.What = strings.Replace(data.What, "occu-", "", -1)
	}
	//
	data.OnlyTimed = (r.PostFormValue("only-timed") == "only-timed")
	//
	data.FirstLineNames = (r.PostFormValue("first-line-names") == "first-line-names")
	//
	data.DateFormat = r.PostFormValue("radio-field-date")
	//
	if r.PostFormValue("radio-sep") == "comma" {
		data.Separator = ","
	} else {
		data.Separator = ";"
	}
	//
	return data, nil
}

