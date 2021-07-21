/******************************************************************************

    Structures used to display a page

    @license    GPL
    @history    2021-07-13 23:51:51+02:00, Thierry Graff : Creation
********************************************************************************/
package ctxt

/**
    Main struct, represented by . in all templates
**/
type Page struct {
	Title   string
	Header  Header
	Footer  Footer
	Details interface{}
}

/**
    Contains fields used in the <head> part of the html page
**/
type Header struct {
	// Value of <title> tag
	Title string
	// Urls of additional css files to load in the header
	CSSFiles []string
	// Urls of js files to load in the header
	// Done before dom loading - most js should be included in footer
	// only for script used by dom construction
	JSFiles []string
}

/**
   Contains information added just before </body> closing tag
   JSFiles are included first, then JSString
**/
type Footer struct {
	// urls of additional scripts to load
	JSFiles []string
	// js code added, surrounded by a <script> tag
	//JSString template.JS
}
