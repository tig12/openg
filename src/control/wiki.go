/**

    Other wiki pages are handled by ShowStaticPage()

**/
package control

import (
	"github.com/gorilla/mux"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	//"fmt"
)


type detailsWiki struct {
	Projects []*model.WikiProject
}

func ShowWiki(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	projects, err := model.GetActiveWikiProjects(ctx.Config.RestURL)
	if err != nil {
		return err
	}
    
	//
	ctx.TemplateName = "wiki.html"
	//
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Wiki",
		},
		Details: detailsWiki{
			Projects: projects,
		},
		Footer: ctxt.Footer{},
	}
	return nil
}


type detailsWikiProject struct {
	Project *model.WikiProject
}

func ShowWikiProject(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	slug := vars["slug"]
	project, err := model.GetWikiProjectFromSlug(ctx.Config.RestURL, slug)
	if err != nil {
		return err
	}
	err = project.ComputePersons(ctx.Config.RestURL)
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
		Details: detailsWikiProject{
			Project: project,
		},
		Footer: ctxt.Footer{},
	}
	return nil
}
