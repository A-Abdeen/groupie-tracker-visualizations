package gt

import (
	"fmt"
	API "gt/webapp/API"
	"html/template"
	"net/http"
	"strconv"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DetailsHandler is called.") // XXX
	// Verify Request Method
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Pattern (Path)
	if r.URL.Path != "/details/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	idParam := r.URL.Query().Get("id")
	idNumber, _ := strconv.Atoi(idParam)
	idNumber--
	var DisplayDetails API.Artists
	DisplayDetails = APIcall[idNumber]
	DisplayDetails.Relations = API.Relations(idNumber)
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		fmt.Println(err.Error()) // XXX
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "details.html", DisplayDetails)
}

