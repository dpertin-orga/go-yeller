package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
	go_version "github.com/hashicorp/go-version"
)

var (
	version string
)

const changelogPath = "CHANGELOG.md"

func getAppVersion() string {
	fileIO, err := os.OpenFile(changelogPath, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\n")
	versionLine := lines[0]
	parsedVersion, err := go_version.NewVersion(strings.Fields(versionLine)[1])
	if err != nil {
		panic(err)
	}
	return parsedVersion.String()
}

func YellingFormat(word string, w io.Writer) {
	fmt.Fprintf(w, "<pre>")
	figure.Write(w, figure.NewFigure("!!! "+strings.ToUpper(word)+" !!!", "puffy", true))
	fmt.Fprintf(w, "</pre>")
	fmt.Fprint(w, "\n<p>"+
		"Powered by <a href=\"https://github.com/dpertin-orga/go-yeller.git\">"+
		"go-yeller</a> v"+version+
		"</p>",
	)
}

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
	YellingFormat(word, w)

}

func main() {
	if len(version) == 0 {
		version = os.Getenv("VERSION")
		if len(version) == 0 {
			version = getAppVersion()
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
