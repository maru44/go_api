package main

import (
	"go_api/handler"
	"go_api/middleware"
	"go_api/sample"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handle(middleware.GetOnlyMiddleware, sample.SampleView))
	http.ListenAndServe(":8080", nil)
}
