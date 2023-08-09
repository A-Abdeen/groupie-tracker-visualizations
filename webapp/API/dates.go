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

func Dates (idNumber int) {
idNumber--
type TmpLocations struct {
    Index []struct {
        Dates []string `json:"dates"`
    } `json:"index"`
}
fullJso, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
if err != nil {
	fmt.Print(err.Error())
	os.Exit(1)
}
fullDatespage, err := ioutil.ReadAll(fullJso.Body)
if err != nil {
	log.Fatal(err)
}
var individualDates TmpLocations
err2 := json.Unmarshal(fullDatespage, &individualDates)
if err2 != nil {
	fmt.Print(err2)
}
fmt.Println(individualDates.Index[idNumber])
}