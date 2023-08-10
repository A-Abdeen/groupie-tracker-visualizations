package gt

import (
	"fmt"
	"io"
	"net/http"
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
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "This is to display details!\n")
}
