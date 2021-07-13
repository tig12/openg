/******************************************************************************
    
    Loads config.yml

    @license    GPL
    @history    2021-07-13 23:49:03+01:00, Thierry Graff : Creation
********************************************************************************/
package ctxt

import (
	"openg.local/openg/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var config *model.Config

// executed at package loading
func init() {
	y, err := ioutil.ReadFile("../config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(y, &config)
	if err != nil {
		panic(err)
	}
}
