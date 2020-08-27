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
	// fmt.Fprintf(w, `<h1>Austin Spencer's Links</h1>
	// <a href="twitter.com/austinnspencer"target="_blank">Twitter</a>
	// <a href="https://www.instagram.com/austinspencer/?hl=en"target="_blank">Instagram</a>
	// <a href="https://www.linkedin.com/in/austin-spencer-b56a25177/"target="_blank">LinkedIn</a>
	// <a href="http://github.com/abspen1"target="_blank">Github</a>`)
	fmt.Fprintf(w, `<h1>Austin Spencer's Links</h1>`)
	fmt.Fprintf(w, "<p></p>")
	fmt.Fprintf(w, `<a href="http:twitter.com/austinnspencer"target="_blank">Twitter</a>`)
	fmt.Fprintf(w, "<p></p>")
	fmt.Fprintf(w, `<a href="http://www.instagram.com/austinspencer/?hl=en"target="_blank">Instagram</a>`)
	fmt.Fprintf(w, "<p></p>")
	fmt.Fprintf(w, `<a href="https://www.linkedin.com/in/austin-spencer-b56a25177/"target="_blank">LinkedIn</a>`)
	fmt.Fprintf(w, "<p></p>")
	fmt.Fprintf(w, `<a href="http://github.com/abspen1"target="_blank">Github</a>`)
}

func main() {
	fmt.Println("Starting Server")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler) // 127.0.0.1:8000/about
	http.HandleFunc("/links/", htmlHandler)  // Adds html tags
	http.ListenAndServe(":8080", nil)        // Run to local server http://127.0.0.1:8000
}
