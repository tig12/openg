/******************************************************************************
    Controls the display of pages related to information sources

    @license    GPL
    @history    2021-07-18 13:22:26+02:00, Thierry Graff : Creation
********************************************************************************/

package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

type detailsSources struct {
	Sources    []*model.Source
	Slug_Name  map[string]string // association slug => name
}

/** 
    Displays a page containing all sources of the database
**/
func ShowSources(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	sources, err := model.GetSources()
	
	slug_Name := make(map[string]string)
	for _, source := range(sources){
	    slug_Name[source.Slug] = source.Name
	}
	
	if err != nil {
		return err
	}
	ctx.TemplateName = "sources.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Information sources",
		},
		Details: detailsSources{
		    Sources: sources,
		    Slug_Name: slug_Name,
		},
	}
	return nil
}