package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/create", create())
	mux.HandleFunc("/getAll", getAll())
	mux.HandleFunc("/getByName", getByName())
	mux.HandleFunc("/deleteByName/", deleteByName())

	return mux
}
