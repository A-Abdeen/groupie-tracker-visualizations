package gt

import (
	"errors"
	"fmt"
	handlers "gt/webapp/handlers"
	"log"
	"net/http"
)

func Server() {
	// Assigned portal (change if needed)
	port := ":8080"
	// Serve the css directory to make it available for html
	css := http.FileServer(http.Dir("../static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	// Main Request Handlers
	http.HandleFunc("/", handlers.BaseHandler)
	http.HandleFunc("/details", handlers.DetailsHandler)
	// Run the server
	fmt.Println("http://localhost:" + port[1:])
	err := http.ListenAndServe(port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("We closed the server? \n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		log.Fatal(err)
	}
}

// Files needed to for parsing templates
