package gt
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)
func Dates(idNumber int) []string{ 
	fullJso, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fullDatespage, err := io.ReadAll(fullJso.Body)
	if err != nil {
		log.Fatal(err)
	}
	var individualDates TmpDates //struct created to be able to unmarshal the full data for locations from above URL
	err2 := json.Unmarshal(fullDatespage, &individualDates)
	if err2 != nil {
		fmt.Print(err2)
	}
	detailsPageDates := individualDates.Index[idNumber]
	return (detailsPageDates.Dates)
}
