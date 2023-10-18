package gt

import (
	"encoding/json"
	"io"
	"net/http"
)

func Locations(id int) ([]string, error) {
	var locations []string
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return locations, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locations, err
	}
	var concertLocations ConcertLocations //struct created to be able to unmarshal the full data for locations from above URL
	if err := json.Unmarshal(body, &concertLocations); err != nil {
		return locations, err
	}
	// locations created so that to be able to add indivitualLocations to DisplayDetails
	artistConcertLocations := concertLocations.Index[id]
	for _, data := range artistConcertLocations.LocationsDetailed {
		data = LocationFmt(data)
		locations = append(locations, data)
	}
	return locations, nil
}
