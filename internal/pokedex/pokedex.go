package pokedex

type pokemon struct {
	Name string
}

type Pokedex struct {
	Pokemon []pokemon
}

func (p *Pokedex) Add(name string) {
	if name == "" {
		return
	}
	pokemon := pokemon{
		Name: name,
	}
	p.Pokemon = append(p.Pokemon, pokemon)
}
