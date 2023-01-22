/**

**/
package ajax

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"openg.local/openg/ctxt"
	"openg.local/openg/model"
)

// To display the download form
type detailsDownloadForm struct {
	Occupations []*model.Group
}

func Download(ctx *ctxt.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	what := vars["what"]
	firstLineNames := vars["firstLineNames"]
	sep := vars["sep"]
	fieldDate := vars["fieldDate"]
	fmt.Printf("what = %s\nfirstLineNames = %s\nsep = %s\nfieldDate = %s\n", what, firstLineNames, sep, fieldDate)
	return nil
}
