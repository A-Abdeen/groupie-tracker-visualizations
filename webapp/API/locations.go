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

func Locations (idNumber int) {
idNumber--
type TmpLocations struct {
    Index []struct {
        Relation []string `json:"locations"`
		dates string `json:"dates"`
    } `json:"index"`
}
fullJso, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
if err != nil {
	fmt.Print(err.Error())
	os.Exit(1)
}
fullLocationspage, err := ioutil.ReadAll(fullJso.Body)
if err != nil {
	log.Fatal(err)
}
var individualLocations TmpLocations
err2 := json.Unmarshal(fullLocationspage, &individualLocations)
if err2 != nil {
	fmt.Print(err2)
}
fmt.Println(individualLocations.Index[idNumber])
}