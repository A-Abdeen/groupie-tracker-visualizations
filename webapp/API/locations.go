package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Locations(idNumber int) {
	idNumber--

	fullJso, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fullLocationspage, err := io.ReadAll(fullJso.Body)
	if err != nil {
		log.Fatal(err)
	}
	var individualLocations TmpLocations
	err2 := json.Unmarshal(fullLocationspage, &individualLocations)
	if err2 != nil {
		fmt.Print(err2)
	}
	// fmt.Println(individualLocations.Index[idNumber]) // XXX
}
