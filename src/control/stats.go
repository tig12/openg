package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

type detailsStats struct {
    Stats   *model.Stats
    CountryNames map[string]string
}

func ShowStats(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
    
    stats, err := model.GetStats(ctx.Config.RestURL)
	if err != nil {
		return err
	}
    
	ctx.TemplateName = "stats.html"
	ctx.Page = &ctxt.Page{
		Header:  ctxt.Header{},
		Details: detailsStats{
		    Stats: stats,
		    CountryNames: model.CountryCodesNames,
		},
	}
	return nil
}
