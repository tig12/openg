/******************************************************************************
    Controls the display of group pages

    @license    GPL
    @history    2021-07-21 08:03:57+02:00 , Thierry Graff : Creation
********************************************************************************/

package control

import (
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	"strconv"
)

type detailsGroup struct {
	Group              *model.Group
	DownloadBase       string
	GroupSlugNames     map[string]string
	Ids_partial_labels map[string]string
	Pages              []int
	CurrentPage        int
	NextPage           int
	PrevPage           int
}

/**
    Displays the page of a group
**/
func ShowGroup(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	slug := vars["slug"]

	strPage := vars["page"]
	page := 1
	if strPage != "" {
		page, _ = strconv.Atoi(strPage) // error useless as routing imposes [1-9][0-9]*
	}

	group, err := model.GetGroupBySlug(ctx.Config.RestURL, slug, page, ctx.Config.NbPerPage)
	if err != nil {
		return err
	}
	if group == nil {
		Show404(w, r)
		return nil
	}

	groupSlugNames, err := model.GetGroupSlugNames(ctx.Config.RestURL)
	if err != nil {
		return err
	}

	title := group.Name
	if group.Type == model.GROUP_TYPE_OCCU {
		title += "s"
	}

	var pages []int
	for i := 1; i <= group.NPages; i++ {
		pages = append(pages, i)
	}

	nextPage := int(math.Min(float64(group.NPages), float64(page+1)))
	prevPage := int(math.Max(1, float64(page-1)))

	ctx.TemplateName = "group.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: title,
		},
		Details: detailsGroup{
			Group:              group,
			DownloadBase:       ctx.Config.Paths.Downloads,
			GroupSlugNames:     groupSlugNames,
			Ids_partial_labels: model.Ids_partial_labels,
			Pages:              pages,
			CurrentPage:        page,
			PrevPage:           prevPage,
			NextPage:           nextPage,
		},
	}
	return nil
}
