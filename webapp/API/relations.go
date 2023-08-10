package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Relations(idNumber int) {
	idNumber--

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
	err2 := json.Unmarshal(fullRelationspage, &individualRelations)
	if err2 != nil {
		fmt.Print(err2)
	}
	fmt.Println(individualRelations.Index[idNumber])
}
