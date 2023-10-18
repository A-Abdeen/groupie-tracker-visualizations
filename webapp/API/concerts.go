package gt

import (
	"encoding/json"
	"io"
	"net/http"
)

func Concerts(idNumber int) (Concert, error) { // Single call for both GetMapData & ConcertDetails in Details Handler
	var concert Concert

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return concert, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return concert, err
	}
	if err := json.Unmarshal(body, &concert); err != nil {
		return concert, err
	}
	return concert, nil
}

func ConcertDetails(concert Concert, artistID int) []string {
	var fmtConcerts []string
	var concertDetails string
	artistConcerts := concert.Index[artistID]
	for idx, data := range artistConcerts.Relation { // since artistConcerts is a map idx is the location and data is the dates for each location
		var dates string
		for idx, i := range data { // i represents each single date and idx is its index
			if idx == 0 {
				dates = dates + i // the array of dates is converted to a string with the dates seperated by a coma
			} else {
				dates = i + ", " + dates
			}
		} // the following three lines are used to better represent each location
		idx = LocationFmt(idx)
		concertDetails = idx + ": " + dates
		fmtConcerts = append(fmtConcerts, concertDetails) // the dates and locations are appended into a single array of string
		dates = ""
	}
	return fmtConcerts
}
