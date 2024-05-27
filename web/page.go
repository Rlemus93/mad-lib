package web

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	var err error

	templates, err = template.New("").ParseGlob("./views/*.html")
	if err != nil {
		panic(err)
	}
}

func ExeTemp(w http.ResponseWriter, name string, data any) error {
	return templates.ExecuteTemplate(w, name, data)
}
