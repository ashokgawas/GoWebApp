// Package model contains just [mongoClient.go] class which contains functions establish connection
// with mongodb
package model

import (
	"context"
	"fmt"
	"time"

	"github.com/ashokgawas/GoWebApp/views"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo connection variables globally defined
// to set once and use in all functions.
var mongoClient *mongo.Client
var mongoContext context.Context

// Connect takes mongodb connection URI as input and
// returns client, context and cancelFunc mongo related variables required to invoke db.
// error variable is populated for any connection error
func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	// Global assignment
	mongoClient = client
	mongoContext = ctx

	return client, ctx, cancel, err
}

// Ping is invoked to verify connectivity with mongodb and returns nothing.
// If application is not connected, error is returned.
func Ping() error {
	if err := mongoClient.Ping(mongoContext, readpref.Primary()); err != nil {
		return err
	}

	//fmt.Println("Mongo Connected")
	return nil
}

// CreateTodo is a peristence call to add a row to todo collection.
//
// It returns mongo result object for successful insertion or error in case of failure.
func CreateTodo(todo views.MongoTodo) (*mongo.InsertOneResult, error) {
	collection := mongoClient.Database("db1").Collection("todo")

	result, error := collection.InsertOne(mongoContext, todo)
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Println("Create todo result - ", result)
	return result, error
}

// GetAllTodo returns all todo entries present in the db
func GetAllTodo() ([]views.MongoTodo, error) {
	collection := mongoClient.Database("db1").Collection("todo")
	var todos []views.MongoTodo

	cursor, err := collection.Find(mongoContext, views.MongoTodo{})

	if err != nil {
		return nil, err
	}

	if err1 := cursor.All(mongoContext, &todos); err1 != nil {
		return nil, err1
	}

	return todos, nil
}

// GetByName returns todo entries based on matching name provided in the input.
func GetByName(name string) ([]views.MongoTodo, error) {
	collection := mongoClient.Database("db1").Collection("todo")
	var todos []views.MongoTodo

	cursor, error := collection.Find(mongoContext, bson.M{"name": name})

	if error != nil {
		return nil, error
	}

	if err := cursor.All(mongoContext, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

// DeleteByName deletes todo entries for the matching name provided in the input.
func DeleteByName(name string) (*mongo.DeleteResult, error) {
	collection := mongoClient.Database("db1").Collection("todo")

	if result, error := collection.DeleteMany(mongoContext, bson.M{"name": name}); error != nil {
		return nil, error
	} else {
		return result, nil
	}
}

// CloseConnection closes mongodb connection.
//
// It is called from [main] function when the server is stopped.
func CloseConnection(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
