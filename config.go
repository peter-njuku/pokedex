package main

import "github.com/peter-njuku/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	pokedex                 map[string]pokeapi.PokemonResponse
	pokedexFile             string
}
