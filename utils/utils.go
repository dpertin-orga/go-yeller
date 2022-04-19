package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	figure "github.com/common-nighthawk/go-figure"
	go_version "github.com/hashicorp/go-version"
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
		"go-yeller</a> v"+getAppVersion()+
		"</p>",
	)
}
