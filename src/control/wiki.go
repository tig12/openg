/**
    This file gathers controllers for "quasi static pages"
**/
package control

import (
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	"net/http"
	"github.com/gorilla/mux"
"fmt"
)

type detailsWikiProject struct {
	Project                   *model.WikiProject
}


func ShowWiki(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	//
	ctx.TemplateName = "wiki.html"
	//
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Wiki",
		},
		Footer: ctxt.Footer{},
	}
	return nil
}

func ShowWikiProject(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	slug := vars["slug"]
	project, err := model.GetWikiProject(ctx.Config.RestURL, slug)
fmt.Printf("%+v\n",project)
	if err != nil {
		return err
	}
	//
	ctx.TemplateName = "wiki-project.html"
	//
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Wiki project: " + project.Name,
		},
		Details: detailsWikiProject {
		    Project: project,
		},
		Footer: ctxt.Footer{},
	}
	return nil
}
