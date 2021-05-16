package http

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeed struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeed(
	httpClient http.Client,
) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	defer close(a.updates)
	// repeatedly:
	for {
		// - get odds from HTTP server
		httpResponse, err := a.httpClient.Get(axilisFeedURL)
		if err != nil {
			return errors.New("error getting axilisFeedURL")
		}

		bodyContent, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return errors.New("error reading http response")
		}

		var offers []axilisOfferOdd
		err = json.Unmarshal(bodyContent, &offers)
		if err != nil {
			return errors.New("error parsing json")
		}

		// - write them to updates channel
		for _, offerOdd := range offers {
			a.updates <- models.Odd{
				Id:          offerOdd.Id,
				Name:        offerOdd.Name,
				Match:       offerOdd.Match,
				Coefficient: 0,
				Timestamp:   time.Time{},
			}
		}

		// - if context is finished, exit and close updates channel
		select {
			case <-ctx.Done():
				return nil

			case <-time.After(time.Second*3):
		}
	}
	// (test your program from cmd/main.go)
	return nil
}

func (a *AxilisOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
}

type axilisOfferOdd struct {
	Id      string
	Name    string
	Match   string
	Details axilisOfferOddDetails
}

type axilisOfferOddDetails struct {
	Price float64
}
