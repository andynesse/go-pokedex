package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/andynesse/go-pokedex/internal/config"
)

func commandExplore(config *config.Config) error {
	if len(config.Args) == 0 {
		return fmt.Errorf("no location to explore")
	}
	fmt.Printf("Exploring %s...\n", config.Args[0])
	poke_api := "https://pokeapi.co/api/v2/location-area/" + config.Args[0]

	res, err := http.Get(poke_api)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var area struct {
		Names []struct {
			Language struct {
				Name string `json:"name"`
			} `json:"language"`
			Name string `json:"name"`
		} `json:"names"`
		PokemonEncounters []struct {
			Pokemon struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"pokemon"`
		} `json:"pokemon_encounters"`
	}
	if err := json.Unmarshal(data, &area); err != nil {
		return err
	}
	fmt.Printf("Found Pokemon in %s:\n", area.Names[0].Name)
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
