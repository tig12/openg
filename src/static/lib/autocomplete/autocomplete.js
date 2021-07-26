/******************************************************************************
        Adaptation from https://www.w3schools.com/howto/howto_js_autocomplete.asp

        @license        GPL
        @history        2020-01-17 16:35:45+01:00, Thierry Graff : Creation
********************************************************************************/

/**
    Generic function which adds html and css to a input type=text field.
    @param inputField     The input text field element on which autocomplete is added
    @param dataProvider   Promise returning a regular array when resolved.
                          Provides the data to use to fill autocomplete values
    Note : le test commenté permet d'avoir "LE PINEL" lorsqu'on tape "PI"
        Le test est utile dans l'exemple de w3school car arr contient tout.
        Inutile ici car arr contient seulement des entrées correspondant à la saisie.
**/
function autocomplete(inputField, dataProvider) {
    var currentFocus;
    inputField.addEventListener(
        "input",
        function(e) {
            let a, b, i, val = this.value;
            closeAllLists();
            // 2 could be passed in param (minimal length before triggering autocomplete)
            if (!val || val.length < 2) {
                return false;
            }
            currentFocus = -1;
            a = document.createElement("DIV");
            a.setAttribute("id", this.id + "autocomplete-list");
            a.setAttribute("class", "autocomplete-items");
            this.parentNode.appendChild(a);
            // here uses function passed in parameter to retrieve values (in general with ajax)
            dataProvider(inputField)
            .then(
                function(response) {
                    for (i = 0; i < response.length; i++) {
                        // voir note
                        //if (response[i].substr(0, val.length).toUpperCase() == val.toUpperCase()) {
                        b = document.createElement("DIV");
                        // version originale, qui met en gras les N premiers caractères
                        //b.innerHTML = "<strong>" + response[i].substr(0, val.length) + "</strong>";
                        //b.innerHTML += response[i].substr(val.length);
                        // version modifiée, qui met en gras le texte saisi ; pas forcément les 2 premiers caractères
                        // Attention, bug si on utilise strong à la place de b :
                        // si val="st" ou "tr" ou "ro" etc. le 2e replace va remplacer à l'intérieur de <strong>
                        const replace = "<b>" + find + "</b>"
                        b.innerHTML = response[i]
                                .replace(val.toUpperCase(), "<b>" + val.toUpperCase() + "</b>")
                                .replace(val.toLowerCase(), "<b>" + val.toLowerCase() + "</b>");
                                // TODO Gérer Mairie qd val = ma (1ere lettre uppercase puis lowercase)
                        b.innerHTML += "<input type='hidden' value='" + response[i] + "'>";
                                b.addEventListener("click", function(e) {
                                inputField.value = this.getElementsByTagName("input")[0].value;
                                closeAllLists();
                        });
                        a.appendChild(b);
                        //}
                }
            }, // end resolve dataProvider promise
            function(error) {
                //alert("Problème autocomplete\nContacter l'administrateur du site");
            }) // end then
            .catch(
                () => { alert("Problème autocomplete\nContacter l'administrateur du site"); }
            );
        } // end event handler
    );
    
    inputField.addEventListener(
        "keydown",
        function(e) {
            var x = document.getElementById(this.id + "autocomplete-list");
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
        var x = document.getElementsByClassName("autocomplete-items");
        for (let i = 0; i < x.length; i++) {
            if (elmnt != x[i] && elmnt != inputField) {
                x[i].parentNode.removeChild(x[i]);
            }
        }
    }

    document.addEventListener(
        "click",
        function (e) {
            closeAllLists(e.target);
        }
    );
} // end function autocomplete
