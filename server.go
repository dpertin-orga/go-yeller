package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func yellingHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}
	word := query.Get("word")
	if len(word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing word")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strings.ToUpper(word)+"§§§")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8888"
	}

	fmt.Println("Yelling server starting on port: " + port)

	http.HandleFunc("/yelling", yellingHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
