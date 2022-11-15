/**

**/
package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	"fmt"
)

// To display the download form
type detailsDownloadForm struct {
	Occupations []*model.Group
}

// To process the download form
type downloadFormData struct {
    What            string
    OnlyTimed       bool
    // output format
    NamesFirstLine  bool
    Separator       string
    DateFormat      string
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
fmt.Printf("%+v\n",data)
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
                Title: "Download | Open Gauquelin database",
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


func processDownloadForm(r *http.Request) (*downloadFormData, error) {
	data := &downloadFormData{}
	var err error
	if err = r.ParseForm(); err != nil {
		return data, err
	}
fmt.Printf("%+v\n",data)
	//
	return data, nil
}

