package dao

import (
	"context"
	"fmt"
	"log"
	"time"
	"user-management-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUser(id string) models.User {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://surya04153:******@mongocluster.7zkgi.mongodb.net/?retryWrites=true&w=majority&appName=MongoCluster").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Access the database and collection
	collection := client.Database("sample_mflix").Collection("users")
	fmt.Println("Collection sample_mflix connected!")
	// ID filter
	idFilter, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": idFilter}

	// Find user with matching ID
	var user models.User
	err = collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found with the specified ID.")
		} else {
			log.Fatal("Error finding document:", err)
		}
	} else {
		// Print the result
		fmt.Printf("User found: ID: %s, Name: %s, Email: %s", user.ID.Hex(), user.Name, user.Email)
	}

	return user

}
