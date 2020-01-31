package bloghandler

import (
	"net/http"

	"html/template"
	"tblog/pkg/terror"
)

type EditHandler struct{}

func (EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/blog/edit.html")
	if err != nil {
		terror.Print(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		terror.Print(err)
	}
}
