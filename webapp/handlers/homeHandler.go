package gt

import (
	"fmt"
	"io"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomeHandler is called.")
	// Verify Request Method
	if r.Method != "GET" {
		// Error Handler
	}
	// Verify Request Path
	if r.URL.Path != "/" {
		// Error Handler
	}
	io.WriteString(w, "This is my website!\n")
}
