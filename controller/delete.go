package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ashokgawas/GoWebApp/model"
)

// deleteByName is package local and deletes todo entries from db
// for the matching "name" passed in the input.
//
// "name" is passed as API path param
func deleteByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			name := strings.TrimPrefix(r.URL.Path, "/deleteByName/")
			fmt.Println(name)
			if deleteResult, error := model.DeleteByName(name); error != nil {
				fmt.Println(error.Error())
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Error Occurred!!"))
			} else {
				fmt.Println("entry deleted -", deleteResult)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				if deleteResult.DeletedCount == 0 {
					json.NewEncoder(w).Encode(struct {
						Status string `json:"status"`
					}{"Nothing Deleted for - " + name})
				} else {
					json.NewEncoder(w).Encode(struct {
						Status string `json:"status"`
					}{"Entry deleted"})
				}

			}

		}
	}
}
