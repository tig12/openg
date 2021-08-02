package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

/** For the page with all occupations **/
type detailsOccus struct {
    Occus   []*model.Occu
}

/** 
    Displays a page listing all occupations
**/
func ShowOccupations(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
    occus, err := model.GetOccus(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	ctx.TemplateName = "occus.html"
	ctx.Page = &ctxt.Page{
		Header:  ctxt.Header{
		    Title: "All occupations",
			CSSFiles: []string{
				"/static/lib/datatables/datatables.min.css"},
			JSFiles: []string{
				"/static/lib/datatables/jquery-3.3.1.min.js",
				"/static/lib/datatables/datatables.min.js"},
		},
		Details: detailsOccus{
		    Occus: occus,
		},
	}
	return nil
}

// ShowOccupation() is hadled by group