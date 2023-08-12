package gt

import (
	"fmt"
	"html/template"
	"net/http"
	gt "gt/webapp/API"
	"strconv"
	// "io"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DetailsHandler is called.")
	// Verify Request Method
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Path
	if r.URL.Path != "/details" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	idNumber,_ := strconv.Atoi(r.FormValue("idNumber"))
	ArtistsDetails := gt.FindArtistFullDetails(idNumber)
	t, err := template.ParseFiles("../templates/details.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	t.Execute(w, ArtistsDetails)
}

