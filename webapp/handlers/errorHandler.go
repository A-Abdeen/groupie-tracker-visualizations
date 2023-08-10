package gt

import (
	"fmt"
	// "io"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	fmt.Println("ErrorHandler is called.")
	var errorMsg string
	switch {
	case statusCode == 400:
		errorMsg = "Bad User Request"
	case statusCode == 404:
		errorMsg = "Page Not Found"
	case statusCode == 405:
		errorMsg = "Method Not Allowed"
	case statusCode == 413:
		errorMsg = "Limit Reached"
	case statusCode == 418:
		errorMsg = "The Server Refuses The Attempt To Brew Coffee With A Teapot"
	case statusCode == 500:
		errorMsg = "Internal Server Error"
	default:
		errorMsg = "UNFAMILIAR ERROR WHAT DIS :("
	}
	errorResponse := Err{false, errorMsg, statusCode}
	http.Error(w, errorResponse.Msg, errorResponse.StatusCode)
	/*
		TODO test how to implement executing templates

		w.WriteHeader(errorResponse.StatusCode)
		io.WriteString(w, "This is to display errors!\n")
	*/

}

// TODO MOVE TO STRUCTS FOLDER LATER
type Err struct {
	IsNil      bool
	Msg        string
	StatusCode int
}
