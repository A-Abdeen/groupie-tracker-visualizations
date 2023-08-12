package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func FindArtistDetails(idNumber int, Whatisrequested string) {
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
	var individualArtists []Artists
	err2 := json.Unmarshal(fullArtistpage, &individualArtists)
	if err2 != nil {
		fmt.Print(err2)
	}
	switch {
	case Whatisrequested == "Id":
		fmt.Println(individualArtists[idNumber].Id)
	case Whatisrequested == "Image":
		fmt.Println(individualArtists[idNumber].Image)
	case Whatisrequested == "Name":
		fmt.Println(individualArtists[idNumber].Name)
	case Whatisrequested == "Member":
		fmt.Println(individualArtists[idNumber].Member)
	case Whatisrequested == "Creationdate":
		fmt.Println(individualArtists[idNumber].Creationdate)
	case Whatisrequested == "FirstAlbum":
		fmt.Println(individualArtists[idNumber].FirstAlbum)
	case Whatisrequested == "Locations":
		Locations(3)
	case Whatisrequested == "Dates":
		Dates((idNumber + 1))
	case Whatisrequested == "Relations":
		Relations((idNumber + 1))
	default:
		fmt.Println("what is requested does not exist")
	}
}
