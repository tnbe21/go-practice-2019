package blog

import (
	"fmt"
	"net/http"

	"tblog/pkg/terror"
	"tblog/pkg/blog/httphandler/blog"
)

func Setup() {
	http.Handle("/", bloghandler.TopHandler{})
	http.Handle("/blog/detail", bloghandler.DetailHandler{})
	http.Handle("/blog/edit", bloghandler.EditHandler{})
	http.Handle("/blog/post", bloghandler.PostHandler{})
	http.Handle("/blog/delete", bloghandler.DeleteHandler{})

	fmt.Println("start server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		terror.Print(err)
	}
}
