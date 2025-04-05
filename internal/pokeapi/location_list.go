package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return RespShallowLocations{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}

	err = json.Unmarshal(data, &locationsResp)

	if err != nil {
		return RespShallowLocations{}, err
	}
	
	fmt.Printf("count: %d, next: %v, previous: %v \n", locationsResp.Count, *locationsResp.Next, *&locationsResp.Previous)
	return locationsResp, nil
}
