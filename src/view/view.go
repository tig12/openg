/******************************************************************************
    File used to embed files of directory view/ in the compiled binary

    @license    GPL
    @history    2021-05-11 15:35:22+02:00, Thierry Graff : Creation
********************************************************************************/
package view

import (
	"embed"
)

//go:embed * admin/*
var ViewFiles embed.FS
