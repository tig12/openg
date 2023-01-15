/******************************************************************************

    Structure to store the contents of config.yml

    @license    GPL
    @history    2021-07-13 23:45:24+02:00, Thierry Graff : Creation
********************************************************************************/
package model

type Config struct {

	// go web server
	Run struct {
		URL  string
		Port string
		Mode string
	}

	// postgrest
	RestURL string `yaml:"rest-url"`

	Paths struct {
		WikiDataDir  string `yaml:"wiki-data-dir"`
		Downloads string `yaml:"download"`
		Log       string
	} // `yaml:"paths"`

	// UI
	NbPerPage int `yaml:"nb-per-page"`
}
