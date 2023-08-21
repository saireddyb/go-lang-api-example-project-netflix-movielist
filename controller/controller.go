package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"netflix-watchlist/model"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const movieCollName = "movies"

var moviesColl *mongo.Collection

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	fmt.Printf("uri: %v\n", uri)
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// validate connect to mongo uri
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Extract the db name from the connection string
	cs, err := connstring.ParseAndValidate(uri)
	if err != nil {
		panic(err)
	}
	dbName := cs.Database

	moviesColl = client.Database(dbName).Collection(movieCollName)
}

func homePage() {
	doc := model.Movies{Name: "KGF", Rating: 5}
	result, err := moviesColl.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	homePage()
	w.Write([]byte("Welcome to the HomePage!"))
}

func getAllMovies() []primitive.M {
	filter := bson.D{}
	cur, err := moviesColl.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error on Finding all the documents", err)
		panic(err)
	}
	defer cur.Close(context.Background())
	var result []primitive.M
	for cur.Next(context.Background()) {
		var doc bson.M
		err := cur.Decode(&doc)
		if err != nil {
			panic(err)
		}
		result = append(result, doc)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	return result
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	result := getAllMovies()
	json.NewEncoder(w).Encode(result)
}

func getMovie(id primitive.ObjectID) primitive.M {
	filter := bson.M{"_id": id}
	var result primitive.M
	err := moviesColl.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document found")
		return result
	}
	if err != nil {
		panic(err)
	}
	return result
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	result := getMovie(id)
	json.NewEncoder(w).Encode(result)
}

func createMovie(movie model.Movies) *mongo.InsertOneResult {
	result, err := moviesColl.InsertOne(context.Background(), movie)
	if err != nil {
		fmt.Println("Error on inserting new Movie", err)
		panic(err)
	}
	return result
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	if movie.Name == "" {
		fmt.Println("name is a required field")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is a required field")
		return
	}
	result := createMovie(movie)
	json.NewEncoder(w).Encode(result)
}
