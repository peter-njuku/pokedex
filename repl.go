package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		args := words[1:]

		availableCommands := getCommands()

		cmd, ok := availableCommands[cmdName]
		if !ok {
			fmt.Println("Unknown command:", cmdName)
			continue
		}
		if err := cmd.callback(cfg, args); err != nil {
			fmt.Println("Error:", err)
		}
	}
}