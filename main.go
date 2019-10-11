package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type shortURL struct {
	ID  string `json:"ID"`
	URL string `json:"URL"`
}

type URLs []shortURL

var shortURLs = URLs{
	{
		ID:  "1",
		URL: "http://jvoljvolizka.xyz",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Kill me")
}

func createURL(w http.ResponseWriter, r *http.Request) {
	var newURL shortURL
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newURL)
	shortURLs = append(shortURLs, newURL)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newURL)
}

func getURL(w http.ResponseWriter, r *http.Request) {
	linkID := mux.Vars(r)["id"]

	for _, oneURL := range shortURLs {
		if oneURL.ID == linkID {
			http.Redirect(w, r, oneURL.URL, 301)
			//json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getURLs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(shortURLs)
}

func deleteURL(w http.ResponseWriter, r *http.Request) {
	linkID := mux.Vars(r)["id"]

	for i, oneURL := range shortURLs {
		if oneURL.ID == linkID {
			shortURLs = append(shortURLs[:i], shortURLs[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", linkID)
		}
	}
}

func main() {
	//initshortURLs()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/shortURL", createURL).Methods("POST")
	router.HandleFunc("/shortURLs", getURLs).Methods("GET")
	router.HandleFunc("/{id}", getURL).Methods("GET")
	router.HandleFunc("/shortURLs/{id}", deleteURL).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
