/******************************************************************************
    Structures describing an event

    @license    GPL
    @history    2021-07-16 23:20:57+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import ()

type Event struct {
	Tzo    string
	Date   string
	Note   string
	Place  Place
	DateUT string `json:"date-ut"`
}

type Place struct {
	C1 interface{} // string
	C2 interface{} // string
	C3 interface{} // string
	Cy string
	//    Lg              float64
	Lg    float32
	//    Lat             float64
	Lat   float32
	Name  string
	Geoid int
}

// ************************** Get fields *******************************

func (p *Place) String() string {
	res := p.Name
	if p.Cy != "" {
		res += " (" + p.Cy + ")"
	}
	return res
}
