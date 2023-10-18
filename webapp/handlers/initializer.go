package gt

import (
	"fmt"
	API "gt/webapp/API"
	"os"
)

/*
The function below is meant to load list of files-
that are parsed for use of html/template library-
at a global scope.
*/

var HtmlTmpl []string // global variables to be used by other functions
var APIcall []API.Artists

func Init() {
	fmt.Println("Initializing Global Variable") // XXX
	HtmlTmpl = []string{
		"../webapp/static/base.html",
		"../webapp/static/details.html",
		"../webapp/static/error.html",
		// Add new html / template names here
	}
	fmt.Println("Global Variable initialized") // XXX
	// used to unmarshal full data into APIcall
	firstCall, err := API.LoadArtist()
	if err != nil {
		fmt.Println("Initial API call failed. Terminating server.")
		os.Exit(0)
	} else {
		APIcall = firstCall
	}
}
