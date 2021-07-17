/******************************************************************************
    Place

    @license    GPL
    @history    2021-07-16 23:20:57+01:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
)


type Place struct{
    C2              string
    C3              string
    Cy              string
//    Lg              float64
    Lg              string
//    Lat             float64
    Lat             string
    Name            string
    Geoid           string
}

// ************************** Get fields *******************************

func (p *Place) String() string {
    res := p.Name
    if p.Cy != "" {
        res += " (" + p.Cy + ")"
    }
	return res
}
