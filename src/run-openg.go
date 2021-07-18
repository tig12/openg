/******************************************************************************
   Entry point of the web application

    @license    GPL
    @history    2021-07-13 22:43:33+02:00, Thierry Graff : Creation
********************************************************************************/
package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
    "io/fs"
	"path/filepath"
	"openg.local/openg/static"
	"openg.local/openg/view"
	"openg.local/openg/ctxt"
	"openg.local/openg/control"
	"openg.local/openg/generic/wilk/werr"
	"github.com/gorilla/mux"
)

// *********************************************************
func main() {

	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("%w", p)
			ctxt.LogError(err)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", H(control.Home))
	r.HandleFunc("/person/{slug:[a-z0-9\\-]+}", H(control.ShowPerson))
	r.HandleFunc("/persons", H(control.ShowPersons))
	r.HandleFunc("/sources", H(control.ShowSources))
	
    // r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(static.StaticFiles))))
	
    // r.PathPrefix("/view/common/").Handler(http.StripPrefix("/view/common/", http.FileServer(http.Dir(filepath.Join("view", "common")))))
    serverView, err := fs.Sub(view.ViewFiles, "view")
    if err != nil {
        log.Fatal(err)
    }
    r.PathPrefix("/view/").Handler(http.StripPrefix("/view/", http.FileServer(http.FS(serverView))))
    //
    serverViewCommon, err := fs.Sub(view.ViewFiles, filepath.Join("view", "common"))
    if err != nil {
        log.Fatal(err)
    }
    r.PathPrefix("/view/common/").Handler(http.StripPrefix("/view/common/", http.FileServer(http.FS(serverViewCommon))))

	r.NotFoundHandler = http.HandlerFunc(notFound)

	ctx := ctxt.NewContext()
	srv := &http.Server{
		Handler:      r,
		Addr:         ctx.Config.Run.URL + ":" + ctx.Config.Run.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// *********************************************************
// H = Handler
// Returns a function with same signature as http.Handler.ServeHTTP() usable by r.HandleFunc()
// Adapter between ServeHTTP() and controller function
// @param  h Controller function
func H(h func(*ctxt.Context, http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := ctxt.NewContext()
		//
		err = h(ctx, w, r) // Call controller h ; fills ctx.TemplateName
		if err != nil {
			showErrorPage(err, ctx, w, r)
			return
		}
		/* 
		if ctx.Redirect != "" {
			http.Redirect(w, r, ctx.Redirect, http.StatusSeeOther)
			return
		}
		*/
		//
		tmpl := ctx.Template
		//
		err = tmpl.ExecuteTemplate(w, "header.html", ctx.Page)
		if err != nil {
			ctxt.LogError(err)
			return
		}
		// Execute template computed by h
		err = tmpl.ExecuteTemplate(w, ctx.TemplateName, ctx.Page)
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
}

// *********************************************************
// Hajax = Handler ajax
// Same as H, but for ajax (does not execute templates)
// @param  h Controller function
func Hajax(h func(*ctxt.Context, http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := ctxt.NewContext()
		err = h(ctx, w, r) // Calls controller h
		if err != nil {
			ctxt.LogError(err)
		}
	}
}

// *********************************************************
// HPDF = Handler PDF
// Same as H, but for pdf (does not execute templates)
// @param  h Controller function
func HPDF(h func(*ctxt.Context, http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := ctxt.NewContext()
		err = h(ctx, w, r) // Calls controller h
		if err != nil {
			ctxt.LogError(err)
		}
	}
}

// *********************** Gestion d'erreur **********************************
// A mettre ailleurs, mais où ?

func notFound(w http.ResponseWriter, r *http.Request) {
	ctx := ctxt.NewContext()
	err := fmt.Errorf("Page inexistante :<br><code><b>%s</b></code>", r.URL)
	showErrorPage(err, ctx, w, r)
}

func showErrorPage(theErr error, ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) {
	type detailsErrorPage struct {
		URL     string
		Details string
	}
	var err error

	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "ERREUR",
		},
		Details: detailsErrorPage{
			URL:     r.URL.String(),
			Details: werr.SprintHTML(theErr),
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
}
