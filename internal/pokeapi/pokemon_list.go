package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(pageUrl *string) (RespSingleLocation, error) {
	url := baseURL + "/location-area/"
	if pageUrl != nil {
		url = url + *pageUrl
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

	return locationResp, nil
}
