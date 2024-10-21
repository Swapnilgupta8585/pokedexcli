package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func PokeapiCall(url string) (Data, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Data{}, fmt.Errorf("can not make a request to the url: %w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Data{}, fmt.Errorf("can not get a response: %w", err)
	}
	if res.StatusCode != 200 {
		return Data{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	var data Data
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return Data{}, fmt.Errorf("can not decode the json: %w", err)
	}
	return data, nil
}
