package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

/** For the page with all occupations **/
type detailsIssues struct {
	Issues             []*model.Issue
	DownloadBase       string
	WD_ENTITY_BASE_URL string
	Slug_Name          map[string]string
}

/**
    Displays a page listing all issues
**/
func ShowIssues(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	issues, err := model.GetIssues(ctx.Config.RestURL)
	if err != nil {
		return err
	}

	ctx.TemplateName = "issues.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Open Gauquelin DB issues",
			CSSFiles: []string{
				"/static/lib/datatables/datatables.min.css"},
			JSFiles: []string{
				"/static/lib/datatables/jquery-3.3.1.min.js",
				"/static/lib/datatables/datatables.min.js"},
		},
		Details: detailsIssues{
			Issues:             issues,
			DownloadBase:       ctx.Config.Paths.Downloads,
			WD_ENTITY_BASE_URL: model.WD_ENTITY_BASE_URL,
		},
	}
	return nil
}

// ShowOccupation() is handled by group
