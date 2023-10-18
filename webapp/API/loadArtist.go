package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func LoadArtist() ([]Artists, error) {
	var artists []Artists //struct created to be able to unmarshal the full data
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return artists, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return artists, err
	}
	if err := json.Unmarshal(body, &artists); err != nil {
		fmt.Print("Hello, World! ") // Indicator of partial marshalling success
	}
	return artists, nil
}
