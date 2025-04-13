package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon (location string) (RespPokemonList, error){
	url := baseURL + "/location-area/" + location


	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespPokemonList{}

		err := json.Unmarshal(val, &pokemonResp)

		if err != nil {
			return RespPokemonList{}, err
		}

		return pokemonResp, nil
	}

	
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
			return RespPokemonList{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
			return RespPokemonList{}, err
	}

	defer res.Body.Close()

	pokemonResp := RespPokemonList{}
	
	data, err := io.ReadAll(res.Body)


	if err != nil {
			return RespPokemonList{}, err
	}

	err = json.Unmarshal(data, &pokemonResp)

	if err != nil {
			return RespPokemonList{}, err
	}

	c.cache.Add(url, data)

	return pokemonResp, nil

}