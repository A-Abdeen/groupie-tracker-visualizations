package gt

import (
	"fmt"
	API "gt/webapp/API"
	"html/template"
	"io"
	"log"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BaseHandler is called.")
	// Verify Request Method
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Path
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "This is my website!\n")
	// TODO Add template execution

	response := API.Artists{}
	t, err := template.ParseFiles("../static/base.html")
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, "base.html", response)
}
