package bloghandler

import (
	"net/http"

	"html/template"
	"tblog/pkg/terror"
)

type DetailHandler struct{}

func (DetailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/blog/detail.html")
	if err != nil {
		terror.Print(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		terror.Print(err)
	}
}
