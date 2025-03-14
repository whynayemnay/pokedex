package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d and body: %s", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	locations := locationStruct{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	if locations.Next != nil {
		cfg.nextUrl = *locations.Next
	} else {
		cfg.nextUrl = ""
	}

	if locations.Previous != nil {
		cfg.previousUrl = *locations.Previous
	} else {
		cfg.previousUrl = ""
	}

	return nil
}

type locationStruct struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
