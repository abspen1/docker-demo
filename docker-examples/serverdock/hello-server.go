package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Austin Spencer!")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Whoa, Go is neat!</h1>
	<p>Go is fast!</p>
	<p>... and simple!</p>
	<p>You %s even add %s</p>
	<a href="http://github.com/abspen1"target="_blank">ME</a>
	`, "can", "<strong>variables</strong>")
	// fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
	// fmt.Fprintf(w, "<p>Go is fast!</p>")
	// fmt.Fprintf(w, "<p>... and simple!</p>")
	// fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler) // 127.0.0.1:8000/about
	http.HandleFunc("/html/", htmlHandler)   // Adds html tags
	http.ListenAndServe(":8080", nil)        // Run to local server http://127.0.0.1:8000
}
