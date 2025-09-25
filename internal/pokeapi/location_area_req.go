package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := BaseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	cache, found := c.cache.Get(fullURL)
	if found {
		LocationAreasResp := LocationAreasResponse{}
		err := json.Unmarshal(cache, &LocationAreasResp)
		if err == nil {
			return LocationAreasResp, nil
		}
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, data)

	LocationAreasResp := LocationAreasResponse{}
	err = json.Unmarshal(data, &LocationAreasResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return LocationAreasResp, nil
}
