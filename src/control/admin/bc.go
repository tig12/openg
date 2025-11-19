/******************************************************************************
    Controls the form to transcribe a birth certificate (bc) to a yaml document.
    To insert or update persons of the database.
    
    Access to this controller is restricted to logged users
    
    @license    GPL
    @history    2025-11-10 20:58:07+01:00, Thierry Graff : Creation
********************************************************************************/

package admin

import (
	"github.com/gorilla/mux"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	"openg.local/openg/control"
	//"fmt"
)

type detailsBCForm struct {
    Subtitle string
    ErrorMessage string
	IsNewPerson bool
	Person model.Person
	HasBC bool
	BCImageURLs []string
}

/**
    Displays the form
**/
func ShowWikiBC(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	
    ctx.TemplateName = "wiki-bc.html"
	vars := mux.Vars(r)
	
	// Variables for ctx.Page.Details
	var isNewPerson bool
	var subtitle string
	var errorMessage string
	var p model.Person
	var hasBC bool
	var bcImageURLs []string
	
	slug, ok := vars["slug"]
    if ok {
        isNewPerson = false
        p, err := model.GetPersonFromSlug(ctx.Config.RestURL, slug)
        if err != nil {
            return err
        }
        if p == nil {
            control.Show404(w, r)
            return nil
        }
        subtitle = p.Name.DisplayedName()
        hasBC = p.HasBC()
        bcImageURLs = model.ComputeBCImageURLs(p, ctx.Config.Paths.WikiDataDir)
    } else {
        isNewPerson = true
        subtitle = "New person"
    }
    //
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Birth certificate transcription",
			CSSFiles: []string{},
		},
		Footer: ctxt.Footer{
			JSFiles: []string{
			},
		},
		Details: detailsBCForm{
			IsNewPerson: isNewPerson,
			Subtitle: subtitle,
			ErrorMessage: errorMessage,
			Person: p,
			HasBC: hasBC,
			BCImageURLs: bcImageURLs,
		},
	}
	return nil
	/* 
	slug := vars["slug"]

	p, err := model.GetPersonFromSlug(ctx.Config.RestURL, slug)
	if err != nil {
		return err
	}
	if p == nil {
		Show404(w, r)
		return nil
	}

	err = p.ComputeGroups(ctx.Config.RestURL)
	if err != nil {
		return err
	}

	err = p.ComputeIssues(ctx.Config.RestURL)
	if err != nil {
		return err
	}

	groupSlugNames, err := model.GetGroupSlugNames(ctx.Config.RestURL)
	if err != nil {
		return err
	}

	sourceSlugNames, err := model.GetSourceSlugNames(ctx.Config.RestURL)
	if err != nil {
		return err
	}

	belongsToHistoricalGroups := false
	for _, g := range p.Groups {
		if g.GroupType == model.GROUP_TYPE_HISTORICAL {
			belongsToHistoricalGroups = true
			break
		}
	}

	hasBC := p.HasBC()
	bcImageURLs := []string{}
	wikiProjects := []*model.WikiProject{}
	if hasBC {
        bcImageURLs = model.ComputeBCImageURLs(p, ctx.Config.Paths.WikiDataDir)
        wikiProjects, err = model.GetWikiProjectsOfBC(ctx.Config.RestURL, p.Acts.Birth)
        if err != nil {
            return err
        }
    }
	ctx.TemplateName = "person.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: p.Name.DisplayedName() + " " + p.GetBirthDay(),
			CSSFiles: []string{
				"/static/lib/tabstrip/tabstrip.css",
				"/static/css/pages/person.css",
			},
		},
		Details: detailsPerson{
			Person:                    p,
			RawFields:                 model.RawPersonSortedFields,
			WikidataEntityURL:         model.WD_ENTITY_BASE_URL,
			Partial_ids_labels:        model.Partial_ids_labels,
			GroupSlugNames:            groupSlugNames,
			SourceSlugNames:           sourceSlugNames,
			BelongsToHistoricalGroups: belongsToHistoricalGroups,
			HasBC:                     hasBC,
			BCImageURLs:               bcImageURLs,
			WikiProjects:              wikiProjects,
		},
		Footer: ctxt.Footer{
			JSFiles: []string{
				"/static/lib/tabstrip/tabstrip.js"},
		},
	}
	return nil
	*/
}
