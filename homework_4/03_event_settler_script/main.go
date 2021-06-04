package main

import (
	"03_event_settler_script/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const host = "127.0.0.1"
const betAPIport = "8787"
const getBetsUrl = "/bets"
const statusKey = "status"

const eventUpdateUrl = "/event/update"
const eventAPIport = "8080"

func getBets() ([]models.BetResponseDto, error) {
	url := fmt.Sprintf("http://%s:%s%s?%s=active", host, betAPIport, getBetsUrl, statusKey)

	resp, err := http.Get(url)
	if err != nil {
		return []models.BetResponseDto{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return []models.BetResponseDto{}, err
	}

	var bets []models.BetResponseDto
	err = json.Unmarshal(body, &bets)
	if err != nil {
		return []models.BetResponseDto{}, err
	}

	return bets, nil
}

func getUniqueSelectionIds(bets []models.BetResponseDto) map[string]bool {
	set := make(map[string]bool)

	for _, b := range bets {
		set[b.SelectionId] = true
	}

	return set
}

func sendEventUpdate(id string, client *http.Client) error {
	url := fmt.Sprintf("http://%s:%s%s", host, eventAPIport, eventUpdateUrl)

	var outcome string
	if rand.Intn(2) == 0 {
		outcome = "won"
	} else {
		outcome = "lost"
	}

	eventDto := models.EventUpdateRequestDto{
		Id:      id,
		Outcome: outcome,
	}

	jsonStr, err := json.Marshal(eventDto)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	rand.Seed(time.Now().Unix())

	bets, err := getBets()
	if err != nil {
		log.Fatal("Error getting bets", err)
	}

	selectionIds := getUniqueSelectionIds(bets)

	client := &http.Client{}
	for id := range selectionIds {
		err = sendEventUpdate(id, client)
		if err != nil {
			log.Fatal("Error sending event update", err)
		}
	}
}
