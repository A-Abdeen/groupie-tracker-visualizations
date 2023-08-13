package gt

import (
	"fmt"
	API "gt/webapp/API"
	"html/template"
	"log"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BaseHandler is called.") // XXX
	// Verify Request Method
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Pattern (Path)
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := API.Artists{} // TODO Pass required details
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, "base.html", response)
}
