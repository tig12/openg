package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

type detailsHome struct {
    Stats   *model.Stats
}

func ShowHome(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
    
    stats, err := model.GetStats(ctx.Config.RestURL)
	if err != nil {
		return err
	}
    
	ctx.TemplateName = "home.html"
	ctx.Page = &ctxt.Page{
		Header:  ctxt.Header{},
		Details: detailsHome{
		    Stats: stats,
		},
	}
	return nil
}
