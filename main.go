package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/peter-njuku/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	result := strings.Fields(text)
	return result
}

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]

		availableCommands := getCommands()

		cmd, ok := availableCommands[cmdName]
		if !ok {
			fmt.Println("Unknown command:", cmdName)
			continue
		}
		if err := cmd.callback(cfg); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewCLient(),
	}

	repl(&cfg)
}
