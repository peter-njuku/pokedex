package pokeapi

import (
	"net/http"
	"time"

	"github.com/peter-njuku/pokedex/internal/pokecache"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewCLient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(5 * time.Second),
	}
}

