package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(pageUrl string) (RespSingleLocation, error) {
	url := baseURL + "/location-area/"
	url = url + pageUrl

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("Cache hit for: ", url)
		location := RespSingleLocation{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return RespSingleLocation{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespSingleLocation{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespSingleLocation{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespSingleLocation{}, err
	}
	locationResp := RespSingleLocation{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespSingleLocation{}, err
	}
	c.cache.Add(url, data)
	return locationResp, nil
}
