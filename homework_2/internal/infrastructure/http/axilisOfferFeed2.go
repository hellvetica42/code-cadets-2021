package http

import (
	"bytes"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
	"context"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const axilisFeedURL2 = "http://18.193.121.232/axilis-feed-2"

type AxilisOfferFeed2 struct {
	updates chan models.Odd
	httpClient *http.Client
}

func NewAxilisOfferFeed2(
	httpClient *http.Client,
) *AxilisOfferFeed2 {
	return &AxilisOfferFeed2{
		updates: make(chan models.Odd),
		httpClient: httpClient,
	}
}

func (a *AxilisOfferFeed2) Start(ctx context.Context) error {
	defer close(a.updates)
	defer log.Printf("shutting down %s", a)

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-time.After(time.Second * 3):
			response, err := a.httpClient.Get(axilisFeedURL2)
			if err != nil {
				log.Println("axilis offer feed 2, http get", err)
				continue
			}
			a.processResponse(ctx, response)
		}
	}
}

func (a *AxilisOfferFeed2) GetUpdates() chan models.Odd {
	return a.updates
}

func (a *AxilisOfferFeed2) String() string {
	return "axilis offer feed 2"
}

func (a *AxilisOfferFeed2) parseResponse(response *http.Response) ([]models.Odd, error) {

	var axilisOffer2Odds []models.Odd

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("axilis offer feed 2, parse response, read all")
		return axilisOffer2Odds, errors.New("could not read response")
	}

	lines := bytes.Split(body, []byte{'\n'})

	for _, line := range lines {
		elements := strings.Split(string(line), ",")
		if len(elements) != 4 {
			log.Println("axilis offer feed 2, parse response, invalid response string")
			continue
		}

		coef, err := strconv.ParseFloat(elements[3], 64)
		if err != nil {
			log.Println("axilis offer feed 2, parse response, invalid coefficient")
			continue
		}

		odd := models.Odd{
			Id: elements[0],
			Name: elements[1],
			Match: elements[2],
			Coefficient: coef,
			Timestamp: time.Now(),
		}

		axilisOffer2Odds = append(axilisOffer2Odds, odd)
	}

	return axilisOffer2Odds, nil

}

func (a *AxilisOfferFeed2) processResponse(ctx context.Context, response *http.Response){
	defer response.Body.Close()

	axilisOffer2Odds, err := a.parseResponse(response)
	if err != nil {
		return
	}

	for _, axilisOdd := range axilisOffer2Odds {
		select {
		case <-ctx.Done():
			return
		case a.updates <- axilisOdd:
			//do nothing
		}
	}
}
