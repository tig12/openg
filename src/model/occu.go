/******************************************************************************
    Occupation

    @license    GPL
    @history    2021-07-17 17:29:31+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
    "sort"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
)

type Occu struct {
	Slug        string
	WD          string
	Name        string
	N           int
	Parents     []string `json:"parents"`
}

// ************************** Get many *******************************

func GetOccus(restURL string) (occus []*Occu, err error) {
	url := restURL + "/occu"
	response, err := http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling "+url)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding occus data")
	}
	if err = json.Unmarshal(responseData, &occus); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal occus data")
	}
	// sort by name
	sorted := make(occuSlice, 0, len(occus))
	for _, elt := range occus {
		sorted = append(sorted, elt)
	}
	sort.Sort(sorted)
	return sorted, nil
}
// Auxiliary of GetOccus() to sort by name
type occuSlice []*Occu
func (o occuSlice) Len() int { return len(o) }
func (o occuSlice) Less(i, j int) bool { return o[i].Name < o[j].Name }
func (o occuSlice) Swap(i, j int) { o[i], o[j] = o[j], o[i] }

