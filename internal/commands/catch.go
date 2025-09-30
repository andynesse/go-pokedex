package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"

	"github.com/andynesse/go-pokedex/internal/config"
)

func commandCatch(config *config.Config) error {
	if len(config.Args) == 0 {
		return fmt.Errorf("no pokemon to catch")
	}
	poke_api := "https://pokeapi.co/api/v2/pokemon/" + config.Args[0]
	res, err := http.Get(poke_api)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var pokemon struct {
		Name           string `json:"name"`
		BaseExperience int    `json:"base_experience"`
	}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	chance := 1.0 / (1.0 + math.Exp((float64(pokemon.BaseExperience)-120)/40.0))
	if chance < rand.Float64() {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	config.Pokedex.Add(pokemon.Name)
	return nil
}
