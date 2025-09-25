package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

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

func (c *Client) GetLocationAreaInfo(areaName string) (LocationAreasExploreResponse, error) {
	areaName = strings.TrimSpace(areaName)
	endpoint := fmt.Sprintf("/location-area/%s", areaName)
	fullURL := BaseURL + endpoint
	cache, found := c.cache.Get(fullURL)
	if found {
		var areaResp LocationAreasExploreResponse
		if err := json.Unmarshal(cache, &areaResp); err == nil {
			return areaResp, nil
		}
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasExploreResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasExploreResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasExploreResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasExploreResponse{}, err
	}

	c.cache.Add(fullURL, data)

	areaResp := LocationAreasExploreResponse{}
	err = json.Unmarshal(data, &areaResp)
	if err != nil {
		return LocationAreasExploreResponse{}, err
	}

	return areaResp, nil
}

func (c *Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	pokemonName = strings.TrimSpace(pokemonName)
	fullURL := BaseURL + "/pokemon/" + pokemonName
	if cached, found := c.cache.Get(fullURL); found {
		var pokemonResp PokemonResponse
		if err := json.Unmarshal(cached, &pokemonResp); err == nil {
			return pokemonResp, nil
		}
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResponse{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, nil
	}

	c.cache.Add(fullURL, data)

	pokemonResp := PokemonResponse{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonResponse{}, nil
	}

	return pokemonResp, nil
}
