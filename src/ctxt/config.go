/******************************************************************************

    Loads config.yml

    @license    GPL
    @history    2021-07-13 23:49:03+02:00, Thierry Graff : Creation
********************************************************************************/
package ctxt

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"openg.local/openg/model"
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
