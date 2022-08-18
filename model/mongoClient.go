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

var mongoClient *mongo.Client
var mongoContext context.Context

func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	// Global assignment
	mongoClient = client
	mongoContext = ctx

	return client, ctx, cancel, err
}

func Ping() error {
	if err := mongoClient.Ping(mongoContext, readpref.Primary()); err != nil {
		return err
	}

	//fmt.Println("Mongo Connected")
	return nil
}

func CreateTodo(todo views.MongoTodo) (*mongo.InsertOneResult, error) {
	collection := mongoClient.Database("db1").Collection("todo")

	result, error := collection.InsertOne(mongoContext, todo)
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Println("Create todo result - ", result)
	return result, error
}

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

func DeleteByName(name string) (*mongo.DeleteResult, error) {
	collection := mongoClient.Database("db1").Collection("todo")

	if result, error := collection.DeleteMany(mongoContext, bson.M{"name": name}); error != nil {
		return nil, error
	} else {
		return result, nil
	}
}

func CloseConnection(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
