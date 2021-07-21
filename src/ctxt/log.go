/******************************************************************************

    Logging

    @license    GPL
    @history    2021-07-13 23:50:21+02:00, Thierry Graff : Creation
********************************************************************************/
package ctxt

import (
	"openg.local/openg/generic/wilk/werr"
)

func LogError(err error) {
	werr.Print(err)
}
