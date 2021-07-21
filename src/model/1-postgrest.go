/******************************************************************************
    Utilities to call posgrest api
    
    Unsuccessful attempt to factorize api calls
    
    url = restURL + "/groop?slug=eq." + slug
    var tmp *interface{}
	tmp, err = GetOneObject(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling GetOneObject("+url+")")
	}
	group = tmp
    
	Piece of operational code:
	// var group *Group
	
    // get the group
	url = restURL + "/groop?slug=eq." + slug
	var body io.Reader
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, werr.Wrapf(err, "Unable to create GET request for " + url)
	}
	request.Header.Set("Accept", "application/vnd.pgrst.object+json") // to have one object returned
    response, err = client.Do(request)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling "+url)
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding group data")
	}
	if err = json.Unmarshal(responseData, &group); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal group data")
	}
    
    @license    GPL
    @history    2021-07-21 09:57:43+02:00, Thierry Graff : Creation
********************************************************************************/
package model

import (
	"encoding/json"
	"io/ioutil"
	"io"
	"net/http"
	"openg.local/openg/generic/wilk/werr"
)

func GetOneObject(url string) (result *interface{}, err error) {
    var responseData []byte
    var response *http.Response
	
	var body io.Reader
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, werr.Wrapf(err, "Unable to create GET request for " + url)
	}
	request.Header.Set("Accept", "application/vnd.pgrst.object+json") // to have one object returned
    response, err = client.Do(request)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling "+url)
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding response body")
	}
	if err = json.Unmarshal(responseData, result); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal")
	}
	return result, nil
}

func GetManyObjects(url string) (result *interface{}, err error) {
    var responseData []byte
    var response *http.Response
    //
	response, err = http.Get(url)
	if err != nil {
		return nil, werr.Wrapf(err, "Error calling "+url)
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, werr.Wrapf(err, "Error decoding response body")
	}
	if err = json.Unmarshal(responseData, result); err != nil {
		return nil, werr.Wrapf(err, "Error json Unmarshal")
	}
	return result, nil
}
