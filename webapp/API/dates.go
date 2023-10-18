package gt

import (
	"encoding/json"
	"io"
	"net/http"
)

func Dates(id int) ([]string, error) {
	var concertsDates []string
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return concertsDates, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return concertsDates, err
	}
	var concertDates ConcertDates
	if err := json.Unmarshal(body, &concertDates); err != nil {
		return concertsDates, err
	}
	artistConcertDates := concertDates.Index[id]
	concertsDates = artistConcertDates.Dates
	return concertsDates, nil
}
