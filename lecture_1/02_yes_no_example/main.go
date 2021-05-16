package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
)

type yesNoResponse struct {
	Answer string
	Forced bool
	Image  string
}

const yesNoURL = "http://yesno.wtf/api"

func linearBackoff(retry int) time.Duration {
	return time.Duration(retry) * time.Second
}

func main() {
	httpClient := pester.New()
	httpClient.Backoff = linearBackoff

	httpResponse, err := httpClient.Get(yesNoURL)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards yesno API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of yesno API response"),
		)
	}

	var decodedContent yesNoResponse
	err = json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	log.Printf("Response from yesno: %v", decodedContent)
}
