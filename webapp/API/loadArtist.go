package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func LoadArtist() []Artists {

	fullJson, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fullArtistpage, err := io.ReadAll(fullJson.Body)
	if err != nil {
		log.Fatal(err)
	}
	var individualArtists []Artists
	err2 := json.Unmarshal(fullArtistpage, &individualArtists)
	if err2 != nil {
		fmt.Print("")
	}
	return individualArtists
}
