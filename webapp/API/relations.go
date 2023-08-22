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
func Relations(idNumber int) []string {
	fullJso, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fullRelationspage, err := io.ReadAll(fullJso.Body)
	if err != nil {
		log.Fatal(err)
	}
	var individualRelations TmpAllConRel
	var dateAndLocation []string
	var oneDateandLocation string
	var dates string
	err2 := json.Unmarshal(fullRelationspage, &individualRelations)
	if err2 != nil {
		fmt.Print(err2)
	}
	detailsPageRelations := individualRelations.Index[idNumber]
	for idx, data := range detailsPageRelations.Relation {
		for idx, i := range data {
			if idx == 0 {
				dates = dates + i
			}else{
			dates = i + ", " + dates
			}
		}
		idx = strings.ReplaceAll(idx, "-", ", ")
		idx = strings.ReplaceAll(idx, "_", " ")
		idx = strings.Title(idx)
 		oneDateandLocation = idx + ": " + dates
		dateAndLocation = append(dateAndLocation, oneDateandLocation)
		dates = ""
	}
	return (dateAndLocation)
}
