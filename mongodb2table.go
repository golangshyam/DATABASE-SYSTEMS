package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if the connection was successful
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to MongoDB")

	// Get a handle to the "users" collection
	usersCollection := client.Database("testing").Collection("users")

	// Get a handle to the "posts" collection
	postsCollection := client.Database("testing").Collection("posts")

	// Insert a user into the "users" collection
	user1 := map[string]interface{}{
		"name":  "shyamvarma",
		"email": "ssk0041@gmail.com",
	}
	userResult, err := usersCollection.InsertOne(context.TODO(), user1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Inserted user:", userResult.InsertedID)

	// Insert a post into the "posts" collection
	post1 := map[string]interface{}{
		"title":     "Hello World",
		"content":   "This is my first post",
		"author_id": userResult.InsertedID,
	}
	postResult, err := postsCollection.InsertOne(context.TODO(), post1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Inserted post:", postResult.InsertedID)

	// Close the connection
	err = client.Disconnect(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connection to MongoDB closed")
}
