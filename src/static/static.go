/******************************************************************************
    File used to embed static files in binary executable

    @license    GPL
    @history    2021-07-14 00:23:06+02:00, Thierry Graff : Creation
********************************************************************************/
package static

import (
    "embed"
)

//go:embed *
var StaticFiles embed.FS
