/**

**/
package control

import (
//"errors"
	"net/http"
	"openg.local/openg/ctxt"
//	"openg.local/openg/model"
)

func ShowDownloads2(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	ctx.TemplateName = "download2.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
            CSSFiles: []string{"/static/css/form.css"},
            //JSFiles: []string{},
		},
	}
	
	/* ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: titrePage,
			JSFiles: []string{
				"/static/js/round.js",
				"/view/common/prix.js"},
		},
		Menu: "chantiers",
		Details: detailsChautreList{
			Chantiers: chantiers,
			Annee:     annee,
			Annees:    annees,
		},
	} */

	return nil
}
