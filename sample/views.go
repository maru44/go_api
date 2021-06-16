package sample

import (
	"errors"
	"go_api/api"
	"net/http"
)

var euro = map[string]string{
	"2008": "Spain",
	"2012": "Spain",
	"2016": "Portugal",
}

func SampleView(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	year := query.Get("year")

	if winner, ok := euro[year]; ok {
		api.JsonResponse(w, map[string]interface{}{"winner": winner})
		return nil
	} else {
		w.WriteHeader(http.StatusNotFound)
		return errors.New("Not Found")
	}
}
