package gt

import (
	// API "gt/webapp/API"
	"html/template"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
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
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "base.html", APIcall) // execution of all artists details to be presented in the homepage using base.html
}
