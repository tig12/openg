/******************************************************************************
    Any document manipulated by the application
    Eg. birth certificate

    @license    GPL
    @history    2023-01-17 23:30:09+01:00, Thierry Graff : Creation
********************************************************************************/
package model

import ()

type DocumentHeader struct {
	History map[string]DocumentHistoryEntry `json:"history"`
}

type DocumentHistoryEntry struct {
	Actor string `json:"actor"`
	Action string `json:"action"`
	Date string `json:"date"`
}

