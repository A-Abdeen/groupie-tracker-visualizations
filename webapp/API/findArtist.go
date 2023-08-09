package gt

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	// "gt/webapp/server"
)

func FindArtist (idNumber int, Whatisrequested string) {
idNumber--
type Artists struct {
	Id int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Member []string `json:"members"`
	Creationdate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
	Locations string `json:"locations"`
	ConcertDates string `json:"concertDates"`
	Relations string `json:"relations"`
}
fullJson, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
if err != nil {
	fmt.Print(err.Error())
	os.Exit(1)
}
fullArtistpage, err := ioutil.ReadAll(fullJson.Body)
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
	Locations((idNumber+1))
case Whatisrequested == "Dates":
	Dates((idNumber+1))
case Whatisrequested == "Relations":
	Relations((idNumber+1))
default:
	fmt.Println("what is requested does not exist")
}
}
