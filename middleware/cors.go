package middleware

import (
	"net/http"
)

func CorsMiddleware(w http.ResponseWriter) error {
	protocol := "http://"
	host := "localhost:3000"
	// こんな感じでローカルかどうか分岐
	// if tools.IsProductionEnv() {
	// 	protocol = "https://"
	// 	host = os.Getenv("FRONT_HOST")
	// }
	w.Header().Set("Access-Control-Allow-Origin", protocol+host)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Cookie")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
	return nil
}
