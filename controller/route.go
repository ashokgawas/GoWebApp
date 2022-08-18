package controller

import (
	"net/http"
)

// Register connects APIs to controller functions to be called upon invocation
// and returns a multiplexer obj
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/create", create())
	mux.HandleFunc("/getAll", getAll())
	mux.HandleFunc("/getByName", getByName())
	// Additional '/' char is required at the end to fetch [name] from path
	mux.HandleFunc("/deleteByName/", deleteByName())

	return mux
}
