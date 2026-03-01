package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonsGetResponse struct {
	Count    int64  `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func main() {
	url := "https://pokeapi.co/api/v2/pokemon"
	fmt.Println("Acessando a API na url", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao acessar a API:", err.Error())
		return
	}
	defer resp.Body.Close()
	fmt.Println("O status code retornado pela url é", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err.Error())
		return
	}

	var pokemonGetResponse PokemonsGetResponse
	err = json.Unmarshal(body, &pokemonGetResponse)
	if err != nil {
		fmt.Println("Erro ao fazer unmarshal do JSON:", err.Error())
		return
	}

	fmt.Println("O nome dos 9 primeiros pokémons é:")
	for i, pokemon := range pokemonGetResponse.Results {
		if i >= 9 {
			break
		}
		fmt.Println(i+1, "-", pokemon.Name)
	}
}
