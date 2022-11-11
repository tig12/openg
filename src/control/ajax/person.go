package ajax

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

/**
    Returns persons whose slug start by a given string (in var "str" of request).
**/
func PersonsAutocomplete(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {

	vars := mux.Vars(r)
	str := vars["str"]
	fmt.Println("=== control/ajax PersonsAutocomplete str = " + str + " ===")

	var err error
	persons, err := model.GetPersonsAutocomplete(ctx.Config.RestURL, str)
	if err != nil {
		return err
	}
	json, err := json.Marshal(persons)
	fmt.Printf("%+v\n", string(json))
	if err != nil {
		return err
	}
	w.Write(json)
	return nil
}
