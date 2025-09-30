package config

import "github.com/andynesse/go-pokedex/internal/pokecache"

type Config struct {
	Next     string
	Previous string
	Cashe    *pokecache.Cache
	Args     []string
}
