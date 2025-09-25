package main

import (
	"encoding/json"
	"os"
)

func savePokedex(cfg *config) error {
	file, err := os.Create(cfg.pokedexFile)
	if err != nil{
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	
	return encoder.Encode(cfg.pokedex)
}
