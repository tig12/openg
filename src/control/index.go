/** 

**/
package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

type detailsHome struct {
	Stats        *model.Stats
	DownloadBase string
}

func ShowHome(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {

	stats, err := model.GetStats(ctx.Config.RestURL)
	if err != nil {
		return err
	}

	ctx.TemplateName = "index.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{},
		Details: detailsHome{
			Stats:        stats,
			DownloadBase: ctx.Config.Paths.Downloads,
		},
	}
	return nil
}

func ShowAbout(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	ctx.TemplateName = "about.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "About the Open Gauquelin Database",
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
