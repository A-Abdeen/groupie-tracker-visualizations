package gt

import (
	GT "gt"
	API "gt/webapp/API"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Verify Request Method
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Pattern (Path)
	if r.URL.Path != "/details/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	// Get the requested Artist ID => int
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	// Validate ID
	if id <= 0 || id > len(APIcall) {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	// Adjust ID to match Index count & load requested artist
	id--
	details := APIcall[id]
	// Call artist's concert details (seperate API calls)
	locationData, err := API.Locations(id)
	if err != nil {
		ErrorHandler(w, r, http.StatusServiceUnavailable)
		return
	}
	details.Locations = locationData
	dateData, err := API.Dates(id)
	if err != nil {
		ErrorHandler(w, r, http.StatusServiceUnavailable)
		return
	}
	details.Dates = dateData
	// Add count for member's list
	details.Member = API.FmtMembers(details.Member)

	concertData, err := API.Concerts(id)
	if err != nil {
		ErrorHandler(w, r, http.StatusServiceUnavailable)
		return
	}
	details.Relations = API.ConcertDetails(concertData, id)
	// GOOGLE MAP DATA // Load and forward Google Map Data: Locations, Coordinates, & Info
	res, err := API.GetMapData(concertData, id)
	if err != nil {
		ErrorHandler(w, r, http.StatusServiceUnavailable)
		return
	}
	details.MapData = res
	// GOOGLE MAP DATA // Load and forward API key from environment
	var apiKey string
	loadKey := func(key string) {
		_, ok := os.LookupEnv(key)
		if !ok {
			GT.LoadAPIKey()
		}
		apiKey = os.Getenv(key)
	}
	loadKey("GOOGLE_MAPS_API_KEY")
	details.MapData.Key = apiKey
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "details.html", details)
}
