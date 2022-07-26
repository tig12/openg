/**

**/
package control

import (
//"errors"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

// Structure containing data for all tabs of home page
type detailsHome struct {
    // common fields
	DownloadBase string
	SelectedTab  string
	// tab intro
	Stats        *model.Stats
	// tab occupations
	Occus              []*model.Group
	WD_ENTITY_BASE_URL string
	Slug_Name          map[string]string
}

func ShowHome(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
    
	selectedTab := ""
	if r.RequestURI == "/" {
		selectedTab = "tab-intro"
	} else if r.RequestURI == "/history" {
		selectedTab = "tab-history"
	} else if r.RequestURI == "/occupations" {
		selectedTab = "tab-occus"
	}
	
    // for tab intro
	stats, err := model.GetStats(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	
	// for tab occupations
	occus, err := model.GetOccus(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	slug_Name, err := model.GetGroupSlugNames(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	
	ctx.TemplateName = "index.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			CSSFiles: []string{
			    "static/css/pages/index.css",
			    "static/css/tabstrip.css",
			},
		},
		Details: detailsHome{
			Stats:        stats,
			DownloadBase: ctx.Config.Paths.Downloads,
			SelectedTab:  selectedTab,
			// for tab occupations
			Occus:              occus,
			WD_ENTITY_BASE_URL: model.WD_ENTITY_BASE_URL,
			Slug_Name:          slug_Name,
		},
		Footer: ctxt.Footer{
			JSFiles: []string{
				"/static/js/tabstrip.js",
			},
		},
	}
	return nil
}

func ShowAbout(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	ctx.TemplateName = "about.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "About | Open Gauquelin DB",
		},
	}
	return nil
}

func ShowFuture(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	ctx.TemplateName = "future.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Future developments | Open Gauquelin DB",
		},
	}
	return nil
}

func ShowDownloads(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	ctx.TemplateName = "download.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{},
	}
	return nil
}

func ShowInstall(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	ctx.TemplateName = "install.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{},
	}
	return nil
}

func ShowWiki(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	ctx.TemplateName = "wiki.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{},
	}
	return nil
}
