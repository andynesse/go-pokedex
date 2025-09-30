package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/andynesse/go-pokedex/internal/config"
	"github.com/andynesse/go-pokedex/internal/pokecache"
)

type area struct {
	Name string
	Url  string
}
type locationResult struct {
	Next     string
	Previous string
	Results  []area
}

func commandMap(config *config.Config) (err error) {
	poke_api := config.Next
	if poke_api == "" {
		return fmt.Errorf("you are already at the last page")
	}
	data, ok := getCachedLocations(config.Cashe, poke_api)
	if !ok {
		data, err = getAPILocations(poke_api)
		if err != nil {
			return err
		}
	}
	locations := locationResult{}
	if err := json.Unmarshal(data, &locations); err != nil {
		return err
	}
	config.Cashe.Add(poke_api, data)
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	config.Next = locations.Next
	config.Previous = locations.Previous
	return nil
}

func commandMapback(config *config.Config) (err error) {
	poke_api := config.Previous
	if poke_api == "" {
		return fmt.Errorf("you are already at the first page")
	}
	data, ok := getCachedLocations(config.Cashe, poke_api)
	if !ok {
		data, err = getAPILocations(poke_api)
		if err != nil {
			return err
		}
	}
	locations := locationResult{}
	if err := json.Unmarshal(data, &locations); err != nil {
		return err
	}
	config.Cashe.Add(poke_api, data)
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	config.Next = locations.Next
	config.Previous = locations.Previous
	return nil
}

func getCachedLocations(cache *pokecache.Cache, poke_api string) ([]byte, bool) {
	val, ok := cache.Get(poke_api)
	if !ok {
		return []byte{}, false
	}
	return val, true
}

func getAPILocations(poke_api string) ([]byte, error) {
	res, err := http.Get(poke_api)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}
