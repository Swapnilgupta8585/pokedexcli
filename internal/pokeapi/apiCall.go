package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Swapnilgupta8585/pokedexcli/internal/pokecache"
)

type Data struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var cache = pokecache.NewCache(60 * time.Second)

func PokeapiCall(url string) (Data, error) {
	value, ok := cache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Data{}, fmt.Errorf("can not make a request to the url: %w", err)
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return Data{}, fmt.Errorf("can not get a response: %w", err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return Data{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}
		dataBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return Data{}, fmt.Errorf("error reading body: %w", err)
		}

		cache.Add(url, dataBytes)

		var data Data
		err = json.Unmarshal(dataBytes, &data)
		if err != nil {
			return Data{}, fmt.Errorf("can not unmarshal the json: %w", err)
		}
		return data, nil
	}

	var data Data
	if err := json.Unmarshal(value, &data); err != nil {
		return Data{}, fmt.Errorf("error unmarshalling json: %w", err)
	}
	return data, nil
}
