package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	collection := MongoClient.Database("articles").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

// create an endpoint to return all articles
func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var articles []Article
	collection := MongoClient.Database("articles").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var article Article
		cursor.Decode(&article)
		articles = append(articles, article)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(articles)
}

// create an endpoint to return a single article
func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := MongoClient.Database("articles").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(article)
}

// create an endpoint to get an article by author
func getArticlesByAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var articles []Article
	params := mux.Vars(r)
	authorID, _ := primitive.ObjectIDFromHex(params["author_id"])
	collection := MongoClient.Database("articles").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{"author_id": authorID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var article Article
		cursor.Decode(&article)
		articles = append(articles, article)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(articles)
}

// create an endpoint to get articles by sport
func getArticlesBySport(w http.ResponseWriter, r *http.Request) {

	log.Println("getArticlesBySport endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	var articles []Article
	params := mux.Vars(r)

	//print params
	log.Println(params)

	sport := params["sport"]

	//check if sport is empty

	if sport == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Sport is required"}`))
		return
	}

	collection := MongoClient.Database("articles").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{"sport": sport})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var article Article
		cursor.Decode(&article)
		articles = append(articles, article)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(articles)
}
