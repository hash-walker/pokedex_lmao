package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client) Pokemon(name string) (RespPokemon, error){

	url := baseURL + "/pokemon/" + name 

	if val, ok := c.cache.Get(url); ok {
		pokemon := RespPokemon{}

		err := json.Unmarshal(val, &pokemon)

		if err != nil {
			return RespPokemon{}, nil
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return RespPokemon{}, nil
	}

	res, err := c.httpClient.Do(req)

	

	if err != nil {
		return RespPokemon{}, nil
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	

	if err != nil {
		return RespPokemon{}, nil
	}

	pokemon := RespPokemon{}

	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return RespPokemon{}, nil
	}

	c.cache.Add(url, data)

	return pokemon, nil

}