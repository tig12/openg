package control

import (
	"net/http"
	"openg.local/openg/ctxt"
)

type detailsHome struct {
}

func Home(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	ctx.TemplateName = "home.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
		},
		Details: detailsHome{
		},
	}
	return nil
}
