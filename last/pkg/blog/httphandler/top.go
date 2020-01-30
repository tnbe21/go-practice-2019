package httphandler

import (
	"net/http"

	"html/template"
	"tblog/pkg/terror"
)

type TopHandler struct{}

func (TopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/blog/index.html")
	if err != nil {
		terror.Print(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		terror.Print(err)
	}
}
