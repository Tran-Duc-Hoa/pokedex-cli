package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchRate := rand.Intn(pokemon.BaseExperience)
	fmt.Println("Catch rate:", catchRate)
	fmt.Println("Base experience:", pokemon.BaseExperience)
	if catchRate < pokemon.BaseExperience / 2 {
		fmt.Printf("Oh no! %s escaped!\n", pokemon.Name)
		return nil
	} 

	fmt.Println("Gotcha!")
	fmt.Printf("You caught a %s!\n", pokemon.Name)
	fmt.Println("Found in the wild")
	cfg.caughtPokemon[pokemon.Name] = pokemon	
	return nil
}
