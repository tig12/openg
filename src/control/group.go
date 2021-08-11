/******************************************************************************
    Controls the display of group pages

    @license    GPL
    @history    2021-07-21 08:03:57+02:00 , Thierry Graff : Creation
********************************************************************************/

package control

import (
	"github.com/gorilla/mux"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	//"fmt"
)

type detailsGroup struct {
	Group        *model.Group
	DownloadBase string
	Slug_Name          map[string]string
}

/**
    Displays the page of a group
**/
func ShowGroup(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	slug := vars["slug"]

	group, err := model.GetGroupBySlug(ctx.Config.RestURL, slug)
	if err != nil {
		return err
	}

	slug_Name, err := model.GetGroupSlugNames(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	
	ctx.TemplateName = "group.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Group " + group.Name,
			CSSFiles: []string{
				"/static/lib/datatables/datatables.min.css"},
			JSFiles: []string{
				"/static/lib/datatables/jquery-3.3.1.min.js",
				"/static/lib/datatables/datatables.min.js"},
		},
		Details: detailsGroup{
			Group:        group,
			DownloadBase: ctx.Config.Paths.Downloads,
			Slug_Name:          slug_Name,
		},
	}
	return nil
}
