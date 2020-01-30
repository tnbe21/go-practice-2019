package blog

import (
	"fmt"
	"net/http"

	"tblog/pkg/terror"
	"tblog/pkg/blog/httphandler"
)

func Setup() {
	http.Handle("/", httphandler.TopHandler{})
	fmt.Println("start server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		terror.Print(err)
	}
}
