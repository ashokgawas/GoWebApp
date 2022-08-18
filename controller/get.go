package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ashokgawas/GoWebApp/model"
	"github.com/ashokgawas/GoWebApp/views"
)

func getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			mongoTodos, err := model.GetAllTodo()

			var todos []views.Todo

			if err != nil {
				fmt.Println(err.Error())
				w.Write([]byte("Exception occured!!"))
				w.WriteHeader(http.StatusInternalServerError)

				return
			} else {
				for i, v := range mongoTodos {
					fmt.Println(i)
					todo := views.Todo{Name: v.Name, Todo: v.Todo, Status: v.Status}
					todos = append(todos, todo)
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(todos)
			}
		}

	}
}

func getByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		var todos []views.Todo

		if mongoTodos, error := model.GetByName(name); error != nil {
			fmt.Println(error.Error())
			w.Write([]byte("Exception occured!!"))
			w.WriteHeader(http.StatusInternalServerError)

			return
		} else {
			for i, v := range mongoTodos {
				fmt.Println(i)
				todo := views.Todo{Name: v.Name, Todo: v.Todo, Status: v.Status}
				todos = append(todos, todo)
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todos)
		}
	}
}
