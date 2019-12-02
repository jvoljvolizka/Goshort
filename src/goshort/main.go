package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

type shortURL struct {
	ID  string `json:"ID"`
	URL string `json:"URL"`
}

var c = GetClient("localhost:27017")

func GetClient(server string) *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://" + server)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func InsertNewURL(client *mongo.Client, shortURL shortURL) interface{} {
	collection := client.Database("Urlshortener").Collection("URLs")
	insertResult, err := collection.InsertOne(context.TODO(), shortURL)
	if err != nil {
		log.Fatalln("Error on inserting new URL", err)
	}
	return insertResult.InsertedID
}

func ReturnURL(client *mongo.Client, filter bson.M) shortURL {
	var shortURL shortURL
	collection := client.Database("Urlshortener").Collection("URLs")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&shortURL)
	return shortURL
}

func ReturnAllURLs(client *mongo.Client, filter bson.M) []*shortURL {
	var URLs []*shortURL
	collection := client.Database("Urlshortener").Collection("URLs")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var URL shortURL
		err = cur.Decode(&URL)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		URLs = append(URLs, &URL)
	}
	return URLs
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Kill me")
}

func createURL(w http.ResponseWriter, r *http.Request) {
	var newURL shortURL
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Hey give me an ID and URL !!!")
	}

	json.Unmarshal(reqBody, &newURL)

	URL := ReturnURL(c, bson.M{"id": newURL.ID})
	if URL.ID == "" {
		InsertNewURL(c, newURL)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newURL)
	} else {
		fmt.Fprintf(w, "Sorry it's already taken")
	}
}

func getURL(w http.ResponseWriter, r *http.Request) {
	linkID := mux.Vars(r)["id"]
	URL := ReturnURL(c, bson.M{"id": linkID})

	http.Redirect(w, r, URL.URL, 301)

}

func getURLs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ReturnAllURLs(c, bson.M{}))
}

/*func deleteURL(w http.ResponseWriter, r *http.Request) {
	linkID := mux.Vars(r)["id"]

	for i, oneURL := range shortURLs {
		if oneURL.ID == linkID {
			shortURLs = append(shortURLs[:i], shortURLs[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", linkID)
		}
	}
}*/

func main() {
	//initshortURLs()
	if len(os.Args) > 1 {
		arg := os.Args[1]
		c = GetClient(arg)
	}

	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/URL", createURL).Methods("POST")
	router.HandleFunc("/URLs", getURLs).Methods("GET")
	router.HandleFunc("/{id}", getURL).Methods("GET")
	//router.HandleFunc("/URL/{id}", deleteURL).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3300", router))
}
