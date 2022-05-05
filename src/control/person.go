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
}

/**
    Displays the page of a single person
**/
func ShowPerson(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	slug := vars["slug"]

	person, err := model.GetPerson(ctx.Config.RestURL, slug)
	if err != nil {
		return err
	}
	if person == nil {
		Show404(w, r)
		return nil
	}

	err = person.ComputeGroups(ctx.Config.RestURL)
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
	for _, g := range person.Groups {
		if g.GroupType == model.GROUP_TYPE_HISTORICAL {
			belongsToHistoricalGroups = true
			break
		}
	}

    err = model.ComputeBC(person, ctx.Config.Paths.Acts)
	if err != nil {
		return err
	}
	
	ctx.TemplateName = "person.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: person.Name.DisplayedName() + " " + person.GetBirthDay(),
			CSSFiles: []string{
				"/static/lib/tabstrip/tabstrip.css",
				"/static/css/pages/person.css",
			},
		},
		Details: detailsPerson{
			Person:                    person,
			RawFields:                 model.RawPersonSortedFields,
			WikidataEntityURL:         model.WD_ENTITY_BASE_URL,
			Partial_ids_labels:        model.Partial_ids_labels,
			GroupSlugNames:            groupSlugNames,
			SourceSlugNames:           sourceSlugNames,
			CountryCodesNames:         model.CountryCodesNames,
			BelongsToHistoricalGroups: belongsToHistoricalGroups,
		},
		Footer: ctxt.Footer{
			JSFiles: []string{
				"/static/lib/tabstrip/tabstrip.js"},
		},
	}
	return nil
}
