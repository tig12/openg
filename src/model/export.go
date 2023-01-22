/******************************************************************************
    Informations about acts (birth, marriage, death certificates).

    @license    GPL
    @history    2022-11-15 20:47:38+01:00, Thierry Graff : Creation
********************************************************************************/
package model

import ()

// To process the download form
type DownloadFormData struct {
	What      string // slug of the group to download
	OnlyTimed bool
	// output format
	FirstLineNames bool
	Separator      string // ";" or ","
	DateFormat     string // "date-iso" or "date-many-cols"
}

func ExportGroup(downloadFormData *DownloadFormData) (result string, err error) {
	return "toto", nil
}
