package main

import (
	"time"

	"github.com/andynesse/go-pokedex/internal/pokecache"
	"github.com/andynesse/go-pokedex/internal/pokedex"
	"github.com/andynesse/go-pokedex/internal/repl"
	"github.com/andynesse/go-pokedex/pkg/config"
)

func main() {
	config := config.Config{
		Next:    "https://pokeapi.co/api/v2/location-area/",
		Cashe:   pokecache.NewCache(5 * time.Second),
		Pokedex: *pokedex.NewPokedex(),
	}

	repl.Run(&config)
}
