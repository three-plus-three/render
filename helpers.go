// +build go1.6

package render

import (
	"fmt"
	"html/template"
)

// Included helper functions for use when rendering HTML.
var helperFuncs = template.FuncMap{
	"yield": func() (string, error) {
		return "", fmt.Errorf("yield called with no layout defined")
	},
	"partial": func() (string, error) {
		return "", fmt.Errorf("block called with no layout defined")
	},
	"current": func() (string, error) {
		return "", nil
	},
	"htmlattr": func(s string) template.HTMLAttr {
		return template.HTMLAttr(s)
	},
	"html": func(s string) template.HTML {
		return template.HTML(s)
	},
	"js": func(s string) template.JS {
		return template.JS(s)
	},
	"set_src": func(s string) template.Srcset {
		return template.Srcset(s)
	},
	"jsstr": func(s string) template.JSStr {
		return template.JSStr(s)
	},
}
