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
	Sources   []*model.Source
	Slug_Name map[string]string // slug => name
	// source slug => paragraph name
	// a paragraph is diaplayed before the current source
	Paragraphs map[string]string
}

/**
    Displays a page containing all sources of the database
**/
func ShowSources(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	sources, err := model.GetSources(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	orderedSources := orderSources(sources)

	slug_Name, err := model.GetSourceSlugNames(ctx.Config.RestURL)
	if err != nil {
		return err
	}
	
	
	/** 
	    source slug => Paragraph title
	    When the view builds source list,
	    The paragraph title is displayed BEFORE the source.
	    (ex: "Primary sources" is displayed before lerrcp source).
	**/
	var paragraphs = map[string]string{
		"lerrcp":       "Primary sources",
		"cura5":        "Secondary sources",
		"a1-booklet":   "Michel and Françoise Gauquelin",
		"afd1-booklet": "Arno Müller",
		"csicop":       "CSICOP (US skeptics)",
		"3a_sports":    "Suitbert Ertel",
	}

	ctx.TemplateName = "sources.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Information sources",
		},
		Details: detailsSources{
			Sources:    orderedSources,
			Slug_Name:  slug_Name,
			Paragraphs: paragraphs,
		},
	}
	return nil
}

/**
    Returns the sources in an arbitrary order.
    Auxiliary of ShowSources().
**/
func orderSources(sources []*model.Source) []*model.Source {
	order := []string{
		// Primary
		"lerrcp",
		"afd",
		// Secondary
		"cura5",
		"newalch",
		"wd",
		"g5",
		// Gauquelin
		"a1-booklet",
		"a1",
		"a2-booklet",
		"a2",
		"a3-booklet",
		"a3",
		"a4-booklet",
		"a4",
		"a5-booklet",
		"a5",
		"a6-booklet",
		"a6",
		"d6-booklet",
		"d6",
		"d10-booklet",
		"d10",
		"e1-booklet",
		"e1",
		"e3-booklet",
		"e3",
		// Müller
		"afd1-booklet",
		"afd1",
		"afd1-100",
		"afd2-booklet",
		"afd2",
		"afd3-booklet",
		"afd3",
		"afd5-booklet",
		"afd5",
		// CSICOP
		"csicop",
		"csi",
		"si42",
		// Ertel
		"3a_sports",
	}
	var res = []*model.Source{}
	for _, slug := range order {
		for _, s := range sources {
			if s.Slug == slug {
				res = append(res, s)
				break
			}
		}
	}
	return res
}
