/******************************************************************************
    Controls the display of error pages.

    @license    GPL
    @history    2022-01-13 21:18:14+01:00 , Thierry Graff : Creation from code previously located in run-openg.go
********************************************************************************/

package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/generic/wilk/werr"
)

/**
    Shows error page.
    Behaviour differ in dev and prod mode.
**/
func ShowErrorPage(theErr error, ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) {
	type detailsErrorPage struct {
		URL     string
		Details string
		Mode    string
	}
	var err error

	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "ERROR",
		},
		Details: detailsErrorPage{
			URL:     r.URL.String(),
			Details: werr.SprintHTML(theErr),
			Mode:    ctx.Config.Run.Mode,
		},
	}
	tmpl := ctx.Template
	err = tmpl.ExecuteTemplate(w, "header.html", ctx.Page)
	if err != nil {
		ctxt.LogError(err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "error.html", ctx.Page)
	if err != nil {
		ctxt.LogError(err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "footer.html", ctx.Page)
	if err != nil {
		ctxt.LogError(err)
		return
	}
	ctxt.LogError(theErr)
}

/**
    Shows Page not found.
    Same behaviour in mode dev and in mode prod.
**/

func Show404(w http.ResponseWriter, r *http.Request) {
	ctx := ctxt.NewContext()
	type detailsErrorPage struct {
		URL string
	}
	var err error

	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "PAGE NOT FOUND",
		},
		Details: detailsErrorPage{
			URL: r.URL.String(),
		},
	}
	tmpl := ctx.Template
	err = tmpl.ExecuteTemplate(w, "header.html", ctx.Page)
	if err != nil {
		ctxt.LogError(err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "error-404.html", ctx.Page)
	if err != nil {
		ctxt.LogError(err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "footer.html", ctx.Page)
	if err != nil {
		ctxt.LogError(err)
		return
	}
}
