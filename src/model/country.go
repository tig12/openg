/******************************************************************************
    Informations about countries

    @license    GPL
    @history    2021-07-22 09:15:26+02:00, Thierry Graff : Creation
********************************************************************************/
package model

/**
    Map country code => country name
**/
var CountryCodesNames map[string]string

// Executed once at package loading
func init() {
	CountryCodesNames = map[string]string{
		"AT": "Austria",
		"BE": "Belgium",
		"CH": "Switzerland",
		"CL": "Chile",
		"CZ": "Czech Republic",
		"DE": "Germany",
		"DK": "Denmark",
		"DZ": "Algeria",
		"ES": "Spain",
		"FR": "France",
		"GB": "United Kingdom",
		"GF": "French Guyana",
		"GP": "Guadeloupe",
		"IT": "Italy",
		"LU": "Luxembourg",
		"MA": "Morroco",
		"MC": "Monaco",
		"MQ": "Martinique",
		"MU": "Mauritius",
		"NL": "Netherlands",
		"PL": "Poland",
		"RU": "Russia",
		"TN": "Tunisia",
		"SE": "Sweden",
		"US": "United States",
	}
}
