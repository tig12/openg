/******************************************************************************
   Entry point of the web application

    @license    GPL
    @history    2021-07-13 22:43:33+02:00, Thierry Graff : Creation
********************************************************************************/
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"openg.local/openg/control"
	"openg.local/openg/control/ajax"
	"openg.local/openg/ctxt"
	"openg.local/openg/static"
	"openg.local/openg/view"
	"path/filepath"
	"time"
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
	
	// ajax
	r.HandleFunc("/ajax/autocomplete/persons/{str}", H(ajax.PersonsAutocomplete))
	
	// routes handled by controls/index.go
	r.HandleFunc("/", H(control.ShowHome))
	r.HandleFunc("/about", H(control.ShowAbout))
	r.HandleFunc("/downloads", H(control.ShowDownloads))
	r.HandleFunc("/future", H(control.ShowFuture))
	r.HandleFunc("/install", H(control.ShowInstall))
	r.HandleFunc("/wiki", H(control.ShowWiki))

	r.HandleFunc("/group/{slug:[a-z0-9\\-]+}", H(control.ShowGroup))
	r.HandleFunc("/group/{slug:[a-z0-9\\-]+}/{page:[1-9][0-9]*}", H(control.ShowGroup))

	r.HandleFunc("/issues", H(control.ShowIssues))
	r.HandleFunc("/issues/{page:[1-9][0-9]*}", H(control.ShowIssues))

	r.HandleFunc("/stats", H(control.ShowStats))
	r.HandleFunc("/sources", H(control.ShowSources))

	r.HandleFunc("/occupations", H(control.ShowOccupations))

	r.HandleFunc("/person/{slug:[a-z0-9\\-]+}", H(control.ShowPerson))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(static.StaticFiles))))

	r.PathPrefix("/view/").Handler(http.StripPrefix("/view/", http.FileServer(http.FS(view.ViewFiles))))

	r.HandleFunc("/download", HDownloadIndex)
	r.PathPrefix("/download/").Handler(http.StripPrefix("/download/", http.FileServer(http.Dir(filepath.Join("..", "download")))))

	r.NotFoundHandler = http.HandlerFunc(control.Show404)

	ctx := ctxt.NewContext()

	r.Use(ctxt.LogRequest) // middleware

	srv := &http.Server{
		Handler:      r,
		Addr:         ctx.Config.Run.URL + ":" + ctx.Config.Run.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// Particular case for one single file
func HDownloadIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("..", "download", "index.html"))
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
			control.ShowErrorPage(err, ctx, w, r)
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
