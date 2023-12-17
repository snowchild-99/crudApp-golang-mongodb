package controller

import (
	"context"
	usermodel "crud-app/userModel"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

//connect with mongodb

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("DB_CONNECT")
	databaseName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")

	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err.Error())

	}
	log.Println("Connection Success with MongoDB Database", databaseName)

	//will return collection
	collection = client.Database(databaseName).Collection(collectionName)
	//collection reference
	log.Println("Collection Name", collection.Name())

}

//mongo helper --insert record

func insertUserInDatabase(users usermodel.Users) {
	insertOne, err := collection.InsertOne(context.Background(), users)

	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Data Inserted Success", insertOne.InsertedID)

}

func deleteUserFromDatabase(userId string) {
	log.Println(userId)
	filter := bson.M{"userid": userId}
	deleteOptions, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("user Deleted", deleteOptions.DeletedCount)

}

func updateUserInDatabase(userId string) {
	log.Println(userId)
	filter := bson.M{"userid": userId}
	update := bson.M{"$set": bson.M{"name": true, "age": true, "email": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Modified Count", result.ModifiedCount)

}

//getAll User from Databse
func getAllUserFromDatabase() []primitive.M {

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err.Error())
	}

	var users []primitive.M
	for cursor.Next(context.Background()) {
		var user bson.M
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(user)
		users = append(users, user)
	}
	defer cursor.Close(context.Background())

	return users

}

func GetAllUser(responseWriter http.ResponseWriter, request *http.Request) {
	//it is fundamental to write header as below but you can modify this as well
	//responseWriter.Header().Set("Content-Type", "sandeep/json")
	responseWriter.Header().Set("Content-Type", "application/json")
	getAllUsers := getAllUserFromDatabase()
	json.NewEncoder(responseWriter).Encode(getAllUsers)

}

//create user
func CreateUser(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	//creating new variable of struct Users
	var user usermodel.Users

	//Decoding the json from requets body into our Users struct object
	error := json.NewDecoder(request.Body).Decode(&user)
	if error != nil {
		log.Println(error.Error())

	}
	log.Println(user)
	insertUserInDatabase(user)
	json.NewEncoder(responseWriter).Encode(user)

}

func UpdateUser(responseWriter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	updateUserInDatabase(params["userid"])

	json.NewEncoder(responseWriter).Encode(params["userid"])

}

func DeleteUser(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	log.Println("Delete User", params["userid"])
	deleteUserFromDatabase(params["userid"])
	json.NewEncoder(responseWriter).Encode(params["userid"])

}
