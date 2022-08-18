package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ashokgawas/GoWebApp/views"
)

// ping is package local and is a health check function to check API connectivity.
func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
