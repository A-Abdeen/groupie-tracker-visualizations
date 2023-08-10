package gt

import (
	"fmt"
	"io"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	fmt.Println("ErrorHandler is called.")
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
