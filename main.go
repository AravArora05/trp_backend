package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// * denotes a pointer

var MongoClient *mongo.Client

// always gets ran before main method, so this is our initialization method
func init() {
	err := godotenv.Load()
	if err != nil {

		log.Printf("Error loading .env file")

	}

	MONGO_URI := os.Getenv("MONGO_URI")

	MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(MONGO_URI))

	if err != nil {
		log.Printf("Error connecting to MongoDB: %v", err)
		panic(err)
	}

	err = MongoClient.Ping(context.Background(), nil)

	if err != nil {

		log.Printf("Error pinging MongoDB: %v", err)
		panic(err)
	}

	log.Println("Connected to MongoDB")

}

func main() {

	/*

		router := http.NewServeMux()



		router.HandleFunc("POST /create-article", createArticle)
		router.HandleFunc("GET /get-articles", getArticles)
		router.HandleFunc("GET /get-article/{'}", getArticle)
		router.HandleFunc("GET /get-by-sport/{sport}", getArticlesBySport)
		router.HandleFunc("GET /get-by-author/{author}", getArticlesByAuthor)

	*/

	r := mux.NewRouter()

	r.HandleFunc("/create-article", createArticle).Methods("POST")

	r.HandleFunc("/get-articles", getArticles).Methods("GET")

	r.HandleFunc("/get-article/{id}", getArticle).Methods("GET")

	r.HandleFunc("/get-by-sport/{sport}", getArticlesBySport).Methods("GET")

	r.HandleFunc("/get-by-author/{author_id}", getArticlesByAuthor).Methods("GET")

	log.Printf("Server started on port 8080")

	http.ListenAndServe(":8080", r)

}
