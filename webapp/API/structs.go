package gt

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Member       []string `json:"members"`
	Creationdate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    []string
	Locations    []string
	Dates        []string
	MapData      MapData `json:"mapData"`
}

// Concert struct to unmarshal the full data for relations
type Concert struct {
	Index []struct {
		Relation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type ConcertLocations struct {
	Index []struct {
		LocationsDetailed []string `json:"locations"`
		DatesDetailed     string   `json:"dates"`
	} `json:"index"`
}
type ConcertDates struct {
	Index []struct {
		Dates []string `json:"dates"`
	} `json:"index"`
}

/*
Below struct is known as struct composition
(or struct embedding)
Allows using single struct to access all resources
Checkout PassError method below
Next step: adding stuct methods for specific handlers
*/
type WebHandler struct {
	*Artists
	Locations *ConcertLocations
	Dates     *ConcertDates
	Relations *Concert
	*Err
}
type Err struct {
	IsErr      bool
	Msg        string
	StatusCode int
}

func (w WebHandler) PassError(x string, y int) WebHandler {
	errorResponse := WebHandler{
		Err: &Err{
			IsErr:      true,
			Msg:        x,
			StatusCode: y,
		},
	}
	return errorResponse
}

/*
Purpose: Send Map Data to Client-Side
(Google Maps request data and detail)
*/

type MapData struct {
	Locations []Location `json:"location"`
	Key       string     `json:"key"`
}

type Location struct {
	Name        string      `json:"name"`
	Coordinates Coordinates `json:"coordinates"`
	Info        string      `json:"info"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

/*
Purpose: Locations => Coordinates
(Rapid API response unmarshal)
*/

type Response struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Bounds struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID string   `json:"place_id"`
		Types   []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}
