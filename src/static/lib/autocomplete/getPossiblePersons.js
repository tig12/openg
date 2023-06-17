/** 
    Promise used by autocomplete() to retrieve data.
    Used by search form.
    @param inputField DOM element (input type=text) used to retrieve person names by ajax.
**/
async function getPossiblePersons(inputField){
    let result = [];
    const url = "/ajax/autocomplete/persons/" + inputField.value;
    let response;
    response = await fetch(url);
    if(response == null){
        alert("ERROR - Transmit this message to the administrator of the site :\n"
            + "Problem to fetch URL: " + url);
        return result;
    }
    response = await response.json();
    response.forEach(function(item) {
        result.push(item);
    });
    return result;
}
