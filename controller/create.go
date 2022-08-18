// Package controller consists of API routes and functions to be invoked when a particular route is called.
//
// [route.go] provides routes to particular functions.
//
// [create.go] contains function to create an entry in the db.
//
// [delete.go] contains function to delete an entry from db based on name.
//
// [get.go] contains 2 functions, to fetch all entries and fetch by name.
//
// [ping.go] contains function for health check i.e. to check mongodb and api connectivity.
package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ashokgawas/GoWebApp/model"
	"github.com/ashokgawas/GoWebApp/views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.Todo{}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				fmt.Println(err.Error())
			}
			//fmt.Println("Received payload -", r.Body)
			fmt.Println("Received payload -", data)
			mongoData := views.MongoTodo{Name: data.Name, Todo: data.Todo, Status: "Open"}
			result, err := model.CreateTodo(mongoData)

			// If error -> 400
			// If success -> 201, with id of mongo inserted row
			if err != nil {
				fmt.Println(err.Error())
				w.Write([]byte("Exception occured!!"))
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusCreated)
				//json.NewEncoder(w).Encode(map[string]string{"id": fmt.Sprint(result.InsertedID)})

				// Better aand cleaner way is to construct dynamic struct
				json.NewEncoder(w).Encode(struct {
					ID string `json:"id"`
				}{fmt.Sprint(result.InsertedID)})
			}
		}
	}
}
