/**

**/
package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

type detailsDownloadForm struct {
	Occupations []*model.Group
}

func ShowDownloads2(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {

	occupations, err := model.GetOccus(ctx.Config.RestURL)
	if err != nil {
		return err
	}

	ctx.TemplateName = "download2.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			CSSFiles: []string{"/static/css/form.css"},
			//JSFiles: []string{},
		},
		Details: detailsDownloadForm{
			Occupations: occupations,
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
	} */

	return nil
}
