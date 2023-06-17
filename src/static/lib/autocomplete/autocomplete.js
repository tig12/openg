/******************************************************************************
    Adaptation of code coming from https://www.w3schools.com/howto/howto_js_autocomplete.asp

    @license        GPL
    @history        2020-01-17 16:35:45+01:00, Thierry Graff : Creation
********************************************************************************/

/**
    Function which adds html and css to a input type=text field.
    @param inputField     The input text field element on which autocomplete is added
    @param dataProvider   Promise returning a regular array when resolved.
                          Each element of this array contains 3 fields: slug, name, day (= birth day)
                          Provides the data to use to fill autocomplete values
**/
function autocomplete(inputField, dataProvider) {
    let currentFocus;
    inputField.addEventListener(
        "input",
        function(e) {
            let a, b, i, val = this.value;
            closeAllLists();
            // 3 could be passed in param (minimal length before triggering autocomplete)
            if (!val || val.length < 3) {
                return false;
            }
            currentFocus = -1;
            a = document.createElement("DIV");
            // here, this = input field => this.id = attribute id of inputField
            a.setAttribute("id", this.id + "autocomplete-list");
            a.setAttribute("class", "autocomplete-items");
            this.parentNode.appendChild(a);
            // here uses function passed in parameter to retrieve values (in general with ajax)
            dataProvider(inputField)
            .then(
                function(response) {
                    for (i = 0; i < response.length; i++) {
                        b = document.createElement("DIV");
                        b.setAttribute('data-slug', response[i].slug);
                        b.innerHTML = response[i].name + ' (' + response[i].day + ')';
                        b.innerHTML += "<input type='hidden' value='" + response[i].name + "'>";
                        b.addEventListener("click", function(e) {
                            // here, this = current div
                            inputField.setAttribute('data-slug', this.getAttribute('data-slug'));
                            inputField.value = this.getElementsByTagName("input")[0].value;
                            closeAllLists();
                        });
                        a.appendChild(b);
                    }
                }, // end resolve dataProvider promise
                function(error) {
                    // alert("Autocomplete problem\nPlease contact site's administrator");
                }
            ) // end then
            .catch(
                () => { alert("Autocomplete problem\nPlease contact site's administrator"); }
            );
        } // end event handler
    );
    
    inputField.addEventListener(
        "keydown",
        function(e) {
            let x = document.getElementById(this.id + "autocomplete-list");
            if (x) x = x.getElementsByTagName("div");
            if (e.keyCode == 40) { // down
                currentFocus++;
                addActive(x);
            } else if (e.keyCode == 38) { // up
                currentFocus--;
                addActive(x);
            } else if (e.keyCode == 27) { // escape
                closeAllLists();
            } else if (e.keyCode == 13) { // enter
                e.preventDefault();
                if (currentFocus > -1) {
                    if (x) x[currentFocus].click();
                }
            }
        }
    );
    
    function addActive(x) {
        if (!x) return false;
        removeActive(x);
        if (currentFocus >= x.length) currentFocus = 0;
        if (currentFocus < 0) currentFocus = (x.length - 1);
        x[currentFocus].classList.add("autocomplete-active");
    }
    
    function removeActive(x) {
        for (let i = 0; i < x.length; i++) {
            x[i].classList.remove("autocomplete-active");
        }
    }
    
    function closeAllLists(elmnt) {
        const x = document.getElementsByClassName("autocomplete-items");
        for (let i = 0; i < x.length; i++) {
            if (elmnt != x[i] && elmnt != inputField) {
                x[i].parentNode.removeChild(x[i]);
            }
        }
    }

    document.addEventListener(
        "click",
        function (e) {
            const slug = inputField.getAttribute('data-slug');
            closeAllLists(e.target);
            if(slug != null){
                document.location.href = '/person/' + slug;
            }
        }
    );
} // end function autocomplete
