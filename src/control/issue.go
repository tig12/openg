package control

import (
    "strconv"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	"github.com/gorilla/mux"
)

/** For the page with all occupations **/
type detailsIssues struct {
	Issues             []*model.Issue
	DownloadBase       string
	WD_ENTITY_BASE_URL string
	GroupSlugNames     map[string]string
	Slug_Name          map[string]string
}

/**
    Displays a page listing all issues
**/
func ShowIssues(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
    
	vars := mux.Vars(r)
	strPage := vars["page"]
	page := 1
	if strPage != "" {
		page, _ = strconv.Atoi(strPage) // error useless as routing imposes [1-9][0-9]*
	}

	issues, err := model.GetIssues(ctx.Config.RestURL, page, ctx.Config.NbPerPage)
	
	if err != nil {
		return err
	}

	// to show the occus of a person
	groupSlugNames, err := model.GetGroupSlugNames(ctx.Config.RestURL)
	if err != nil {
		return err
	}

	ctx.TemplateName = "issues.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Persons' issues",
		},
		Details: detailsIssues{
			Issues:             issues,
			DownloadBase:       ctx.Config.Paths.Downloads,
			GroupSlugNames:     groupSlugNames,
			WD_ENTITY_BASE_URL: model.WD_ENTITY_BASE_URL,
		},
	}
	return nil
}
