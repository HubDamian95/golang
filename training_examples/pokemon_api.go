package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type PokemonAPIResponse struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	// var anyData map[string]any
	var myResponse PokemonAPIResponse
	json.Unmarshal(responseData, &myResponse)
	// fmt.Println(myResponse["name"])

	// fmt.Println(myResponse.Name)
	// fmt.Println(myResponse.Pokemon)

	for i := 0; i < len(myResponse.Pokemon); i++ {
		fmt.Println(myResponse.Pokemon[i])
	}

}
