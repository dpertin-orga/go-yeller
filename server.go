package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/dpertin-orga/go-yeller/utils"
)

var (
	version string
)

const changelogPath = "CHANGELOG.md"

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
	utils.YellingFormat(word, version, w)

}

func main() {
	if len(version) == 0 {
		version = os.Getenv("VERSION")
		if len(version) == 0 {
			version = utils.GetAppVersion(changelogPath)
		}
	}
	fmt.Println("Version: " + version)

	port := ":" + os.Getenv("PORT")
	if len(port) == 1 {
		port = ":8888"
	}

	fmt.Println("Yelling server starting on port: " + port)

	http.HandleFunc("/yelling", yellingHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
