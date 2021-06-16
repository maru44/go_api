package main

import (
	"go_api/handler"
	"go_api/sample"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handle(sample.SampleView))
	http.ListenAndServe(":8080", nil)
}
