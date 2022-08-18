// main hosts the server for to-do API application.
//
// To-Do application contains APIs providing these features:
//   - Create new entry
//   - Fetch all entries from the db
//   - Fetch entries from db based on name
//   - Delete entries based on name
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashokgawas/GoWebApp/controller"
	"github.com/ashokgawas/GoWebApp/model"
)

func main() {
	client, ctx, cancel, err := model.Connect("mongodb://myadmin:myadmin@localhost:27017/admin")
	if err != nil {
		panic(err)
	}

	defer model.CloseConnection(client, ctx, cancel)

	// Ping mongodb and check connectivity
	if err = model.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Mongodb connected")

	mux := controller.Register()
	log.Fatal(http.ListenAndServe(":8000", mux))

}
