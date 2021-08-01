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

/** For the page of one particular occupation **/
type detailsOccu struct {
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
		Header:  ctxt.Header{},
		Details: detailsOccus{
		    Occus: occus,
		},
	}
	return nil
}

/** 
    Displays a page listing one occupation
**/
func ShowOccupation(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
    
	ctx.TemplateName = "occu.html"
	ctx.Page = &ctxt.Page{
		Header:  ctxt.Header{},
		Details: detailsOccu{
		},
	}
	return nil
}
