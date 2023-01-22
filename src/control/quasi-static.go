/**
    This file gathers controllers for "quasi static pages"
**/
package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

// Home
type detailsHome struct {
	Stats *model.Stats
}

func ShowHome(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	stats, err := model.GetStats(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	ctx.TemplateName = "home.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			CSSFiles: []string{
				"static/css/pages/index.css",
			},
		},
		Details: detailsHome{
			Stats: stats,
		},
		Footer: ctxt.Footer{},
	}
	return nil
}

// Historical datasets
type detailsHistory struct {
	DownloadBase string
}

func ShowHistory(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	//
	ctx.TemplateName = "history.html"
	//
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Historical datasets",
		},
		Details: detailsHistory{
			DownloadBase: ctx.Config.Paths.Downloads,
		},
		Footer: ctxt.Footer{},
	}
	return nil
}

// Lists by profession
type detailsOccus struct {
	Occus              []*model.Group
	DownloadBase       string
	WD_ENTITY_BASE_URL string
	Slug_Name          map[string]string
}

func ShowOccupations(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	//
	ctx.TemplateName = "occus.html"
	//
	occus, err := model.GetOccus(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	slug_Name, err := model.GetGroupSlugNames(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Professional groups",
		},
		Details: detailsOccus{
			Occus:              occus,
			DownloadBase:       ctx.Config.Paths.Downloads,
			WD_ENTITY_BASE_URL: model.WD_ENTITY_BASE_URL,
			Slug_Name:          slug_Name,
		},
		Footer: ctxt.Footer{},
	}
	return nil
}

func ShowCandidates(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	//
	ctx.TemplateName = "data-candidates.html"
	//
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Data candidates",
		},
		Footer: ctxt.Footer{},
	}
	return nil
}

func ShowAbout(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	//
	ctx.TemplateName = "about.html"
	//
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "About OGBD",
		},
		Footer: ctxt.Footer{},
	}
	return nil
}

// TODO remove this page
func ShowFuture(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	//
	ctx.TemplateName = "future.html"
	//
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Future developments",
		},
		Footer: ctxt.Footer{},
	}
	return nil
}

func ShowInstall(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	//
	ctx.TemplateName = "install.html"
	//
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Installation",
		},
		Footer: ctxt.Footer{},
	}
	return nil
}

// TODO REMOVE when download2 is finished
type detailsDownload struct {
	DownloadBase string
}

func ShowDownloads(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	//
	ctx.TemplateName = "download.html"
	//
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Downloads",
		},
		Details: detailsDownload{
			DownloadBase: ctx.Config.Paths.Downloads,
		},
		Footer: ctxt.Footer{},
	}
	return nil
}
