/******************************************************************************
    Context, contains variables usable by all pages

    @license    GPL
    @history    2021-07-13 23:37:37+02:00, Thierry Graff : Creation
********************************************************************************/
package ctxt

import (
	"html/template"
	"openg.local/openg/model"
)

type Context struct {
	Page         *Page
	Redirect     string
	TemplateName string
	Template     *template.Template
	Config       *model.Config
}

func NewContext() *Context {
	ctx := &Context{}
	ctx.Template = tmpl // declared in template.go
	ctx.Config = config // declared in config.go - see also model/1-config.go
	return ctx
}
