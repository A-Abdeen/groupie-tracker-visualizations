package gt

import (
	"log"
	"regexp"
	"strings"
)

func Slug(input string) string {
	var output string

	regxCh, err1 := regexp.Compile(`[^\w]`)
	if err1 != nil {
		log.Fatal(err1)
	}
	regxRS, err2 := regexp.Compile(`\s+`)
	if err2 != nil {
		log.Fatal(err2)
	}
	output = strings.ReplaceAll(input, "@", "at")
	output = strings.ReplaceAll(output, "&", "and")
	output = strings.ToLower(output)
	output = strings.ReplaceAll(output, "$", "S")
	output = regxCh.ReplaceAllString(output, " ")
	output = regxRS.ReplaceAllString(output, " ")
	output = strings.ReplaceAll(output, " ", "-")
	if string(output[0]) == "-" {
		output = output[1:]
	}
	if string(output[len(output)-1]) == "-" {
		output = output[:len(output)-1]
	}
	return output
}
