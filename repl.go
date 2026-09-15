package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/itency/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	commands         map[string]cliCommand
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		args := []string{}
		if len(cleanInput(input)) > 1 {
			args = cleanInput(input)[1:]
		}
		cmd, ok := cfg.commands[cleanInput(input)[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg, args...); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func cleanInput(text string) []string {
	cleanText := []string{}
	textSplit := strings.Split(text, " ")
	for _, word := range textSplit {
		loweredWord := strings.ToLower(word)
		if loweredWord != "" {
			cleanText = append(cleanText, loweredWord)
		}
	}
	return cleanText
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the Pokemon found in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays information about a caught Pokemon",
			callback:    commandInspect,
		},
	}
}
