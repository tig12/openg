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

func ShowPerson(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	slug := vars["slug"]
    
	person, err := model.GetPerson(slug)
	if err != nil {
		return err
	}
	ctx.TemplateName = "person.html"
	ctx.Page = &ctxt.Page{
		Header: ctxt.Header{
			Title: "Person",
		},
		Details: detailsPerson{
		    Person: person,
		},
	}
	return nil
}
