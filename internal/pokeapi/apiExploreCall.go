package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Swapnilgupta8585/pokedexcli/internal/pokecache"
)

type Explore struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

var ExploreCache = pokecache.NewCache(60 * time.Second)

func ExploreApiCall(url string) (Explore, error) {
	value, ok := ExploreCache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Explore{}, fmt.Errorf("can not make a request to the url: %w", err)
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return Explore{}, fmt.Errorf("can not get a response: %w", err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return Explore{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}
		ExploreBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return Explore{}, fmt.Errorf("error reading body: %w", err)
		}

		ExploreCache.Add(url, ExploreBytes)

		var explore Explore
		err = json.Unmarshal(ExploreBytes, &explore)
		if err != nil {
			return Explore{}, fmt.Errorf("can not unmarshal the json: %w", err)
		}
		return explore, nil
	}

	var explore Explore
	if err := json.Unmarshal(value, &explore); err != nil {
		return Explore{}, fmt.Errorf("error unmarshalling json: %w", err)
	}
	return explore, nil
}
