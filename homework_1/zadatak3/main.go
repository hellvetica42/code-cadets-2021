package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
)

type PokemonLocation struct {
	LocationName string `json:"name"`
}

type PokemonEncounter struct {
	LocationArea PokemonLocation `json:"location_area"`
}

type PokemonInfo struct {
	Name string `json:"name"`
	EncountersURL string `json:"location_area_encounters"`
}

type PokemonInfoOutput struct {
	Name string `json:"name"`
	Locations []PokemonLocation `json:"locations"`
}

const mainURL = "https://pokeapi.co/api/v2/pokemon/"

func main() {

	var pokemonNameOrId string

	flag.StringVar(&pokemonNameOrId, "identifier", "", "Name or ID of pokemon")
	flag.Parse()

	if len(pokemonNameOrId) == 0 {
		log.Fatal(errors.New("no identifier input"))
	}

	httpClient := pester.New()

	httpResponse, err := httpClient.Get(mainURL + pokemonNameOrId)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards pokemon API"),
		)
	}

	if httpResponse.StatusCode == 404 {
		log.Fatal(
			errors.New("Pokemon with identifier not found"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)

	var decodedPokemonInfo PokemonInfo
	err = json.Unmarshal(bodyContent, &decodedPokemonInfo)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content for pokemon info"),
		)
	}

	httpResponse, err = httpClient.Get(decodedPokemonInfo.EncountersURL)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards pokemon API"),
		)
	}

	bodyContent, err = ioutil.ReadAll(httpResponse.Body)

	var decodedEncountersData []PokemonEncounter
	err = json.Unmarshal(bodyContent, &decodedEncountersData)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content for pokemon encounters"),
		)
	}

	if len(decodedEncountersData) == 0 {
		log.Fatal(errors.New("Pokemon has no discoverable locations."))
	}

	var locations []PokemonLocation

	for _, encounter := range decodedEncountersData {
		locations = append(locations, encounter.LocationArea)
	}

	var outputStruct  = PokemonInfoOutput {
		Name: decodedPokemonInfo.Name,
		Locations: locations,
	}

	output, err := json.Marshal(outputStruct)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "marshalling the JSON body content for pokemon info output"),
		)
	}

	fmt.Println(string(output))
}
