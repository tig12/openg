/******************************************************************************
    Statistics about the database

    @license    GPL
    @history    2021-07-22 01:01:51+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"openg.local/openg/generic/tiglib"
	"openg.local/openg/generic/wilk/werr"
	//"fmt"
)

type Stats struct {
	N         int
	N_time    int
	N_notime  int
	N_countries int
	N_checked int
	Countries map[string]int
	Years     map[string]int // `json:"years"`
	// not in database
	YearMin string
	YearMax string
}

// ************************** Get one *******************************

func GetStats(restURL string) (stats *Stats, err error) {
	var url string
	var responseData []byte
	var response *http.Response

	// get the stats
	url = restURL + "/stats"
	response, err = http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling postgrest API: \n"+url)
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding Stats data")
	}
	var tmp []*Stats
	if err = json.Unmarshal(responseData, &tmp); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal Stats\n"+string(responseData)+"\n")
	}
	if len(tmp) == 0 {
		var stats = Stats{}
		return &stats, werr.Wrapf(err, "EMPTY Stats - need to be initialized")
	}
	stats = tmp[0]

	// year min, year max
	years := tiglib.MapKeysStringInt(stats.Years)
	stats.YearMin, stats.YearMax = tiglib.MinMaxString(years)
	stats.N_countries = len(stats.Countries)

	return stats, nil
}
