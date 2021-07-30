/** 
    Promise used by autocomplete() to retrieve data
    Used by search form
    
    @param inputField input type=text to retrieve person names by ajax
**/
async function getPossiblePersons(inputField){
    let result = [];
    const url = "/ajax/autocomplete/acteur/" + inputField.value;
    let response;
    response = await fetch(url);
    if(response == null){
        alert("- ERROR - Transmettez this message to the site administrator :\n"
            + "Preoble to fetch " + url);
        return result;
    }
    response = await response.json();
    response.forEach(function(item) {
        result.push(item.name);
    });
    return result;
}
