/******************************************************************************
    File used to embed files of directory static/ in the binary executable

    @license    GPL
    @history    2021-07-14 00:23:06+02:00, Thierry Graff : Creation
********************************************************************************/
package static

import (
	"embed"
)

//go:embed *
var StaticFiles embed.FS
