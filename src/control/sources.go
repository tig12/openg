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
	Slug_Name  map[string]string // slug => name
	// source slug => paragraph name
	// a paragraph is diaplayed before the current source
	Paragraphs map[string]string
}

/** 
    Displays a page containing all sources of the database
**/
func ShowSources(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	sources, err := model.GetSources()
	if err != nil {
		return err
	}
	slug_Name, err := model.GetSourceSlugNames()
	if err != nil {
		return err
	}
	
	var paragraphs = map[string]string {
	    "g5": "Main sources",
	    "a1-booklet": "Michel and Françoise Gauquelin",
	    "afd": "Arno Müller",
	    "csicop-committee": "CSICOP",
	}
	
	ctx.TemplateName = "sources.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Information sources",
		},
		Details: detailsSources{
		    Sources: sources,
		    Slug_Name: slug_Name,
		    Paragraphs: paragraphs,
		},
	}
	return nil
}