package control

import (
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	"strconv"
)

/** For the page with all occupations **/
type detailsIssues struct {
	Issues             []*model.Issue
	DownloadBase       string
	WD_ENTITY_BASE_URL string
	GroupSlugNames     map[string]string
	Slug_Name          map[string]string
	NIssues            int
	Pages              []int
	CurrentPage        int
	NextPage           int
	PrevPage           int
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

	issues, nIssues, pagemax, err := model.GetIssues(ctx.Config.RestURL, page, ctx.Config.NbPerPage)
	if err != nil {
		return err
	}
	var pages []int
	for i := 1; i <= pagemax; i++ {
		pages = append(pages, i)
	}
	nextPage := int(math.Min(float64(pagemax), float64(page+1)))
	prevPage := int(math.Max(1, float64(page-1)))

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
			NIssues:            nIssues,
			Pages:              pages,
			CurrentPage:        page,
			PrevPage:           prevPage,
			NextPage:           nextPage,
		},
	}
	return nil
}
