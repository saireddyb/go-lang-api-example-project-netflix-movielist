package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const directorCollName = "directors"
const movieCollName = "movies"
const countryCollName = "countries"

var DirectorColl *mongo.Collection
var MovieColl *mongo.Collection
var CountryColl *mongo.Collection

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	cs, err := connstring.ParseAndValidate(uri)
	if err != nil {
		panic(err)
	}
	dbName := cs.Database
	DirectorColl = client.Database(dbName).Collection(directorCollName)
	MovieColl = client.Database(dbName).Collection(movieCollName)
	CountryColl = client.Database(dbName).Collection(countryCollName)

}
