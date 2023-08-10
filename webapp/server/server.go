package gt

import (
	"errors"
	"fmt"
	handlers "gt/webapp/handlers"
	"net/http"
	"os"
)

func Server() {
	port := ":8080"
	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("http://localhost:" + port[1:])
	err := http.ListenAndServe(port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("We closed the server? \n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
