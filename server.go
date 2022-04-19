package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/dpertin-orga/go-yeller/utils"
)

func yellingHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	word := query.Get("word")
	if len(word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing word")
		return
	}
	w.WriteHeader(http.StatusOK)
	utils.YellingFormat(word, w)

}

func main() {
	/*
		version := os.Getenv("VERSION")
		if len() == 0 {
			port = ":8888"
		}
	*/

	port := ":" + os.Getenv("PORT")
	if len(port) == 1 {
		port = ":8888"
	}

	fmt.Println("Yelling server starting on port: " + port)

	http.HandleFunc("/yelling", yellingHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
