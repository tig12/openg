package control
/* 

import (
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Autocomplete(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {

	vars := mux.Vars(r)
	str := vars["str"]

	type respElement struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	var resp []respElement

	var err error
	acteurs, err := model.GetActeursAutocomplete(ctx.DB, str)
	if err != nil {
		return err
	}
	var nom string
	for _, a := range acteurs {
		if a.Prenom == "" {
			nom = a.Nom
		} else {
			nom = a.Nom + " " + a.Prenom
		}
		resp = append(resp, respElement{a.Id, nom})

	}
	json, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	w.Write(json)
	return nil
}
*/