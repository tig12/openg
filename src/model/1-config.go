/******************************************************************************
    
    Structure to store the contents of config.yml
    
    @license    GPL
    @history    2021-07-13 23:45:24+01:00, Thierry Graff : Creation
********************************************************************************/
package model

type Config struct {
	Run struct {
		URL  string `yaml:"url"`
		Port string `yaml:"port"`
	}
	Paths struct {
		Acts string `yaml:"acts"`
	} `yaml:"paths"`
}
