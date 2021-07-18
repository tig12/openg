/******************************************************************************
    Controls the display of person pages

    @license    GPL
    @history    2021-07-16, Thierry Graff : Creation
********************************************************************************/

package control

import (
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	"github.com/gorilla/mux"
)

type detailsPerson struct {
	Person    *model.Person
}
type detailsPersons struct {
	Persons    []*model.Person
}

/** 
    Displays the page of a single person
**/
func ShowPerson(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	slug := vars["slug"]
    
	person, err := model.GetPerson(slug)
	if err != nil {
		return err
	}
//person.Name.Official.Family = "test official family"
//person.Name.Official.Given = "test official given"
//person.Name.NobiliaryParticle = "von"

	ctx.TemplateName = "person.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: person.GetName() + " " + person.GetBirthDay(),
		},
		Details: detailsPerson{
		    Person: person,
		},
	}
	return nil
}

/** 
    Displays a page containing a list of persons
**/
func ShowPersons(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	persons, err := model.GetPersons()
	if err != nil {
		return err
	}
	ctx.TemplateName = "persons.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Person list",
			CSSFiles: []string{
				"/static/lib/datatables/datatables.min.css"},
			JSFiles: []string{
			    "/static/lib/datatables/jquery-3.3.1.min.js",
                "/static/lib/datatables/datatables.min.js"},
		},
		Details: detailsPersons{
		    Persons: persons,
		},
	}
	return nil
}