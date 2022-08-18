package views

import "go.mongodb.org/mongo-driver/bson/primitive"

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type Todo struct {
	Name   string `json:"name"`
	Todo   string `json:"todo"`
	Status string `json:"status"`
}

type MongoTodo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Todo   string             `bson:"todo,omitempty"`
	Status string             `bson:"status,omitempty"`
}
