/******************************************************************************
    File used to embed files of directory view/ in the compiled binary

    @license    GPL
    @history    2021-05-11 15:35:22+02:00, Thierry Graff : Creation
********************************************************************************/
package wiki

import (
	"embed"
)

//go:embed *
var ViewWikiFiles embed.FS
