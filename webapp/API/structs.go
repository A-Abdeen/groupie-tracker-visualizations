package gt

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Member       []string `json:"members"`
	Creationdate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type TmpAllConRel struct {
	Index []struct {
		Relation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type TmpLocations struct {
	Index []struct {
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type TmpDates struct {
	Index []struct {
		Dates []string `json:"dates"`
	} `json:"index"`
}
