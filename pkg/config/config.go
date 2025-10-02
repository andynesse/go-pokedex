package config

import (
	"github.com/andynesse/go-pokedex/internal/pokecache"
	"github.com/andynesse/go-pokedex/internal/pokedex"
)

type Config struct {
	Next     string
	Previous string
	Cashe    *pokecache.Cache
	Args     []string
	Pokedex  pokedex.Pokedex
}
