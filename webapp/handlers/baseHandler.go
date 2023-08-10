package gt

import (
	"fmt"
	"io"
	"net/http"
	"gt/webapp/API"
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
	FullArtistsDetails := gt.FindArtistFullDetails(4)
	fmt.Println(FullArtistsDetails)
}
