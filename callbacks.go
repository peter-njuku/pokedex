package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/peter-njuku/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the command line",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the map",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Return a lot more information about the location area",
			callback: func(cfg *config, args []string) error {
				if len(args) != 1 {
					fmt.Println("Usage: explore <loaction-area>")
					return nil
				}
				areaName := strings.TrimSpace(args[0])
				return commandExplore(cfg, areaName)
			},
		},
		"catch": {
			name:        "catch",
			description: "Throws pokeball to pikachu",
			callback: func(cfg *config, args []string) error {
				if len(args) != 1 {
					fmt.Println("Usage: catch <name>")
					return nil
				}
				pokemonName := strings.TrimSpace(args[0])
				return commandCatch(cfg, pokemonName)
			},
		},
		"inspect": {
			name:        "inspect",
			description: "Provides information about a pokemon",
			callback: func(cfg *config, args []string) error {
				if len(args) != 1 {
					fmt.Println("Usage: inspect <name>")
					return nil
				}
				pokemonName := strings.TrimSpace(args[0])
				return commandInspect(cfg, pokemonName)
			},
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the pokemons caught",
			callback:    commandPokedex,
		},
	}
}

func getCaughtPokemonNames(cfg *config) []string {
	var names []string
	if cfg.pokedex == nil {
		return names
	}
	for name := range cfg.pokedex {
		names = append(names, name)
	}
	return names
}

func commandPokedex(cfg *config, args []string) error {
	names := getCaughtPokemonNames(cfg)
	for _, n := range names {
		fmt.Println("- " + n)
	}

	return nil
}

func caught(cfg *config, pokemonName string) bool {
	if cfg.pokedex == nil {
		return false
	}
	_, caught := cfg.pokedex[pokemonName]
	return caught
}

func commandInspect(cfg *config, pokemonName string) error {
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func commandCatch(cfg *config, pokemonName string) error {
	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	catchProb := 1.0 - float64(resp.BaseExperience)/500.0
	if catchProb < 0.05 {
		catchProb = 0.05
	} else if catchProb > 0.95 {
		catchProb = 0.95
	}

	if rand.Float64() <= catchProb {
		fmt.Printf("%s was caught!\n", pokemonName)
		if cfg.pokedex == nil {
			cfg.pokedex = make(map[string]pokeapi.PokemonResponse)
		}
		cfg.pokedex[resp.Species.Name] = resp
		if err := savePokedex(cfg); err != nil {
			fmt.Println("Could not save pokedex")
		} else {
			fmt.Println("Pokedex updated successfully")
		}

		fmt.Println("You may inspect it using the inspect command")
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
		fmt.Printf("You have %d Pokémon in your pokedex.\n", len(cfg.pokedex))
	}
	return nil
}

func commandExplore(cfg *config, areaName string) error {
	resp, err := cfg.pokeapiClient.GetLocationAreaInfo(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", resp.Name)
	var names []string
	fmt.Println("Found Pokemon:")
	for idx, encounters := range resp.PokemonEncounters {
		names = append(names, encounters.Pokemon.Name)
		fmt.Printf("- %s\n", names[idx])
	}
	return nil
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Exiting Pokedex... Goodbye!!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, args []string) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("you are on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}
