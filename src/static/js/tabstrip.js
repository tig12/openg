/******************************************************************************
    
    @license    GPL - conforms to file LICENCE located in root directory of current repository.
    
    @history    2018, Thierry Graff : Creation
********************************************************************************/

document.getElementsByClassName("defaultOpen")[0].click();


function openTab(tabName) {
    let i, tabcontent, tablinks;
    tabcontent = document.getElementsByClassName("tabcontent");
    for (i = 0; i < tabcontent.length; i++) {
        tabcontent[i].style.display = "none";
    }
    tablinks = document.getElementsByClassName("tablinks");
    for (i = 0; i < tablinks.length; i++) {
        tablinks[i].className = tablinks[i].className.replace(" active", "");
    }
    document.getElementById(tabName).style.display = "block";
    // tab contents are named 'tab-intro', 'tab-history', 'tab-occus'
    // buttons are named 'intro', 'history', 'occus'
    const btnName = tabName.substr(4); // remove 'tab-' to get button name
    document.getElementById(btnName).className += " active";
    const tab_url = {
        'intro'         : '',
        'history'       : 'history',
        'occus'         : 'occupations',
        'wiki'          : 'wiki',
        'candidates'    : 'candidates',
    };
    const tab_title = {
        'intro'         : 'Open Gauquelin Database',
        'history'       : 'Historical datasets | Open Gauquelin Database',
        'occus'         : 'Lists by occupations | Open Gauquelin Database',
        'wiki'          : 'Wiki | Open Gauquelin Database',
        'candidates'    : 'Data candidates | Open Gauquelin Database',
    };
//console.log(tab_url[btnName] + ' - ' + tab_title[btnName]);
    history.pushState({}, '', tab_url[btnName]);
    document.title = tab_title[btnName];
}