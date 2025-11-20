package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello World %q", html.EscapeString(r.URL.Path))
		if err != nil {
			log.Fatalln(fmt.Errorf("error writing response: %v", err))
			return
		}
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
		if err != nil {
			log.Fatalln(fmt.Errorf("hi: %w", err))
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8070", nil))

}
