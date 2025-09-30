package pokedex

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  map[string]int
	Types  []string
}

type Pokedex struct {
	Pokemon map[string]Pokemon
}

func NewPokedex() *Pokedex {
	pokedex := Pokedex{
		Pokemon: make(map[string]Pokemon),
	}
	return &pokedex
}

func (p *Pokedex) Add(name string, height, weight int, stats map[string]int, types []string) {
	newPokemon := Pokemon{
		Name:   name,
		Height: height,
		Weight: weight,
		Stats:  stats,
		Types:  types,
	}
	p.Pokemon[name] = newPokemon
}
