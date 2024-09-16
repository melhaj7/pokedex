package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(url string) (*locationResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var locationResponse locationResponse
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return nil, err
	}
	return &locationResponse, nil
}
