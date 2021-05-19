package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"os"
)

type application struct {
	Name string
	Age int
	Passed bool
	Skills []string
}

const url = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func contains(arr []string, item string) bool{

	for _, val := range arr{
		if val == item{
			return true
		}
	}
	return false

}

func main() {
	httpClient := pester.New()

	httpResponse, err:= httpClient.Get(url)

	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards json API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of json API response"),
		)
	}

	var applications []application
	err = json.Unmarshal(bodyContent, &applications)

	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	var valid_names []application
	for _, val := range applications{
		if val.Passed && (contains(val.Skills, "Go") || contains(val.Skills, "Java")){
			valid_names = append(valid_names, val)
		}
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}
	defer f.Close()

	for _, val := range valid_names{
		f.WriteString(fmt.Sprint(val.Name) + " - ")
		for _, skill := range val.Skills{
			f.WriteString(fmt.Sprint(skill) + " ")
		}
		f.WriteString("\n")
	}


}
