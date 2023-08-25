package gt
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)
func Locations(idNumber int) []string{
	fullJso, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fullLocationspage, err := io.ReadAll(fullJso.Body)
	if err != nil {
		log.Fatal(err)
	}
	var individualLocations TmpLocations //struct created to be able to unmarshal the full data for locations from above URL
	err2 := json.Unmarshal(fullLocationspage, &individualLocations)
	if err2 != nil {
		fmt.Print(err2)
	}
	var locations []string // locations created so that to be able to add indivitualLocations to DisplayDetails 
	detailsPageLocations := individualLocations.Index[idNumber]
	for _, data := range detailsPageLocations.LocationsDetailed {
		data = strings.ReplaceAll(data, "-", ", ")
		data = strings.ReplaceAll(data, "_", " ")
		data = strings.Title(data)
		locations = append(locations, data)
	}
	return (locations)
}

