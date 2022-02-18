package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

/** For the page with all occupations **/
type detailsOccus struct {
	Occus              []*model.Group
	DownloadBase       string
	WD_ENTITY_BASE_URL string
	Slug_Name          map[string]string
}

/**
    Displays a page listing all occupations
**/
func ShowOccupations(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	occus, err := model.GetOccus(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	slug_Name, err := model.GetGroupSlugNames(ctx.Config.RestURL)
	if err != nil {
		return err
	}

	ctx.TemplateName = "occus.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Lists by occupation",
			CSSFiles: []string{
				"/static/lib/datatables/datatables.min.css",
			    "/static/css/pages/occus.css"},
			JSFiles: []string{
				"/static/lib/datatables/jquery-3.3.1.min.js",
				"/static/lib/datatables/datatables.min.js"},
		},
		Details: detailsOccus{
			Occus:              occus,
			DownloadBase:       ctx.Config.Paths.Downloads,
			WD_ENTITY_BASE_URL: model.WD_ENTITY_BASE_URL,
			Slug_Name:          slug_Name,
		},
	}
	return nil
}

// ShowOccupation() is handled by group
