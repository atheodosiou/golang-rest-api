package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title       string `json:"Title"`
	Descripiton string `json:"Description`
	Content     string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title 1", Descripiton: "Test Description 1", Content: "Hello World 1"},
		Article{Title: "Test Title 2", Descripiton: "Test Description 2", Content: "Hello World 2"},
		Article{Title: "Test Title 3", Descripiton: "Test Description 3", Content: "Hello World 3"},
		Article{Title: "Test Title 4", Descripiton: "Test Description 4", Content: "Hello World 4"},
		Article{Title: "Test Title 5", Descripiton: "Test Description 5", Content: "Hello World 5"},
		Article{Title: "Test Title 6", Descripiton: "Test Description 6", Content: "Hello World 6"},
		Article{Title: "Test Title 7", Descripiton: "Test Description 7", Content: "Hello World 7"},
		Article{Title: "Test Title 8", Descripiton: "Test Description 8", Content: "Hello World 8"},
		Article{Title: "Test Title 8", Descripiton: "Test Description 9", Content: "Hello World 9"},
		Article{Title: "Test Title 10", Descripiton: "Test Description 10", Content: "Hello World 10"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

// func newArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "New article endpoint works")
// }

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simle REST API using Golang and gorilla/mux.")
	fmt.Fprintf(w, "Try a get request at /articles...")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	// myRouter.HandleFunc("/articles", newArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequest()
}
