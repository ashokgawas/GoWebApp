// Package views contains structure types used in the application.
package views

import "go.mongodb.org/mongo-driver/bson/primitive"

// A Response returned with /ping API
//	-	Code: int value representing HTTP Response Code
//	-	Body: type value representing description relating to the response code.
type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

// A Todo type for I/O of data from/to client in json format.
type Todo struct {
	Name   string `json:"name"`
	Todo   string `json:"todo"`
	Status string `json:"status"`
}

// A MongoTodo type to pass data to/from mongodb (specifically)
type MongoTodo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Todo   string             `bson:"todo,omitempty"`
	Status string             `bson:"status,omitempty"`
}
