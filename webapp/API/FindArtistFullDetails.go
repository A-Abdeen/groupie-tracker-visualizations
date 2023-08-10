package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func FindArtistFullDetails(idNumber int) Artists{
	idNumber--

	fullJson, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fullArtistpage, err := io.ReadAll(fullJson.Body)
	if err != nil {
		log.Fatal(err)
	}
	var AllArtistsDetails []Artists
	err2 := json.Unmarshal(fullArtistpage, &AllArtistsDetails)
	if err2 != nil {
		fmt.Print(err2)
	}
	
	return AllArtistsDetails[idNumber]
}