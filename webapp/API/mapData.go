package gt

import (
	"encoding/json"
	"fmt"
	GT "gt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetMapData(concert Concert, artistID int) (MapData, error) {
	var mapData MapData
	var newLocation Location
	var newCords Coordinates
	concertsList := concert.Index[artistID]
	for location, dateArray := range concertsList.Relation { // since concertDetails is a map, key is location name and dateArray is list of dates for each location
		var dates []string
		dates = append(dates, dateArray...)
		// the following three lines are used to better represent each location
		res, err := RapidApi(location)
		if err != nil {
			return mapData, err
		}
		newCords = res
		newLocation.Coordinates = newCords
		location = LocationFmt(location)
		newLocation.Name = location
		newLocation.Info = InfoFmt(location, dates)
		mapData.Locations = append(mapData.Locations, newLocation)
	}
	return mapData, nil
}

func RapidApi(input string) (Coordinates, error) {
	var output Coordinates
	var apiKey string
	loadKey := func(key string) {
		_, ok := os.LookupEnv(key)
		if !ok {
			GT.LoadAPIKey()
		}
		apiKey = os.Getenv(key)
	}
	url := fmt.Sprint("https://google-maps-geocoding.p.rapidapi.com/geocode/json?address=", input, "&language=en")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return output, err
	}
	loadKey("X-RapidAPI-Key")
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "google-maps-geocoding.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return output, err
	}
	defer res.Body.Close() // BUG add to all API requests
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return output, err
	}
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return output, err
	}
	for _, lvl1 := range response.Results {
		output.Lat = lvl1.Geometry.Location.Lat
		output.Lng = lvl1.Geometry.Location.Lng
	}
	return output, nil
}

func LocationFmt(location string) string {
	location = strings.ReplaceAll(location, "-", ", ")
	location = strings.ReplaceAll(location, "_", " ")
	location = strings.Title(location)
	return location
}

func InfoFmt(label string, dates []string) string {
	var info string
	info = fmt.Sprintln("\n<h2>" + label + "</h2>")
	for _, date := range dates {
		info = fmt.Sprintln(info + "<p>" + date + "</p>")
	}
	return info
}
