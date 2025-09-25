package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/peter-njuku/pokedex/internal/pokeapi"
)

func loadPokedex(cfg *config) error {
	file, err := os.Open(cfg.pokedexFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Creating new  pokedex storage....")
			cfg.pokedex = make(map[string]pokeapi.PokemonResponse)
			return savePokedex(cfg)
		}
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err = decoder.Decode(&cfg.pokedex); err != nil {
		return err
	}

	fmt.Printf("Loaded %d pokemon from pokedex.\n", len(cfg.pokedex))
	return nil
}
