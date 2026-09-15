package main

import (
	"time"

	"github.com/itency/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeapiClient,
	}
	startRepl(cfg)
}
