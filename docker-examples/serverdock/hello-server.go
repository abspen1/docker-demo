package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
*	Simple RESTful API created with GOlang
*	This is using localhost:8080
*
 */

// Article is struct for article JSON
// type Article struct {
// 	Title   string `json:"title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// Articles is a slice of Article
// type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	// articles := Articles{
	// 	Article{Title: "Test Ttile", Desc: "Test Description", Content: "Hello World"},
	// }
	fmt.Println("Endpoint Hit: Are you sure you want ALL articles?")
	// json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Whoa, Go is neat!</h1>
	<p>Go is fast!</p>
	<p>... and simple!</p>
	<p>You %s even add %s</p>
	<a href="{{ github.com/abspen1 }}" target="_blank">{{ ME }}</a>
	`, "can", "<strong>variables</strong>")
	// fmt.Fprintf(w, "Test POST endpoint worked!")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Austin Spencer")
	// fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
