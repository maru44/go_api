package middleware

import (
	"errors"
	"net/http"
)

// preflightに対してstatus okを返す
// そうしないとPUTやDELETEのような非単純リクエストが実行されない
func AllowOptionsMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	return nil
}

func UpsertOnlyMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" || r.Method == "PUT" {
		return nil
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return errors.New("METHOD NOT ALLOWED")
}

func PostOnlyMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return nil
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return errors.New("METHOD NOT ALLOWED")
}

func PutOnlyMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "PUT" {
		return nil
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return errors.New("METHOD NOT ALLOWED")
}

func DeleteOnlyMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "DELETE" {
		return nil
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return errors.New("METHOD NOT ALLOWED")
}

func GetOnlyMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return nil
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return errors.New("METHOD NOT ALLOWED")
}
