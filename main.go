package main

import (
	"fmt"

	"github.com/peter-njuku/pokedex/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewCLient(),
		pokedex:       make(map[string]pokeapi.PokemonResponse),
		pokedexFile:   "pokedex.json",
	}

	if err := loadPokedex(&cfg); err != nil {
		fmt.Println("Failed to load Pokedex:", err)
	}

	repl(&cfg)
}
