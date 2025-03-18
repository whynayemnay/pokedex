package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) EncounterPokemon(pokemonName string) (RespPokemonEncounter, error) {
	url := baseURL + "/pokemon/"
	url = url + pokemonName

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("Cache hit for: ", url)
		pokemon := RespPokemonEncounter{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return RespPokemonEncounter{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonEncounter{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonEncounter{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonEncounter{}, err
	}
	pokemonResp := RespPokemonEncounter{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespPokemonEncounter{}, err
	}
	c.cache.Add(url, data)
	return pokemonResp, nil
}
