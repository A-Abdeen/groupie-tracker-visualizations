package gt

import (
	"fmt"
	API "gt/webapp/API"
	"html/template"
	"log"
	"net/http"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DetailsHandler is called.") // XXX
	// Verify Request Method
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Pattern (Path)
	if r.URL.Path != "/details" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := API.Artists{} // TODO Pass required details
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, "details.html", response)
}
