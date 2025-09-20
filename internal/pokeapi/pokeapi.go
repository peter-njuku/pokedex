package pokeapi

import (
	"net/http"
	"time"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func NewCLient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
