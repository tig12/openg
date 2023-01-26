/******************************************************************************
    Controls the display of person pages

    @license    GPL
    @history    2021-07-16, Thierry Graff : Creation
********************************************************************************/

package control

import (
	"github.com/gorilla/mux"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
"fmt"
)

type detailsPerson struct {
	Person                    *model.Person
	RawFields                 map[string][]string
	Partial_ids_labels        map[string]string
	WikidataEntityURL         string
	GroupSlugNames            map[string]string
	SourceSlugNames           map[string]string
	CountryCodesNames         map[string]string
	BelongsToHistoricalGroups bool
	HasBC                     bool // HasBirthCertificate
	BCImageURLs               []string
	WikiProjects              []*model.WikiProject
}

/**
    Displays the page of a single person
**/
func ShowPerson(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
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
        wikiProjects, err = model.ComputeBCWikiProjects(ctx.Config.RestURL, p.Acts.Birth)
        if err != nil {
            return err
        }
    }
fmt.Printf("wikiProjects = %+v\n",wikiProjects)

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
			CountryCodesNames:         model.CountryCodesNames,
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
}
