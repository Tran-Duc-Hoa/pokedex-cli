package main

import (
	"errors"
	"fmt"
	"strings"
)



func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	
	name := args[0]
	pokemon, exists := cfg.caughtPokemon[strings.ToLower(name)]

	if !exists {
		fmt.Printf("You have not caught %s\n", name)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, value := range pokemon.Types {
		fmt.Printf("  -%s\n", value.Type.Name)
	}
	return nil
}