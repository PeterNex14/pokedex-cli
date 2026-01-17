package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/PeterNex14/pokedex-cli/internal/pokecache"
)

type ListLocation struct {
	Next		*string 	`json:"next"`
	Previous 	*string		`json:"previous"`
	Results		[]struct {
		Name	string	`json:"name"`
		URL		string	`json:"url"`
	}	`json:"results"`
}

type PokemonLocation struct {
	PokemonEncounter		[]struct {
		Pokemon			struct {
			Name		string			`json:"name"`
			Url 		string			`json:"url"`
		}	`json:"pokemon"`
	}	`json:"pokemon_encounters"`
}

type Pokemon struct {
	Name			string		`json:"name"`
	BaseExperience 	int			`json:"base_experience"`
}



func getLocationArea(baseUrl string, c *pokecache.Cache) (ListLocation, error) {
	reqCache, ok := c.Get(baseUrl)
	var listLocation ListLocation

	if ok {
		err := json.Unmarshal(reqCache, &listLocation)
		if err != nil {
			return listLocation, err
		}

		return listLocation, nil
	}
	
	res, err := http.Get(baseUrl)
	if err != nil {
		return listLocation,  err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return listLocation, err
	}

	if err := json.Unmarshal(data, &listLocation); err != nil {
		return listLocation, err
	}

	c.Add(baseUrl, data)
	return listLocation, nil

}


func getPokemonLocation(location, baseUrl string, c *pokecache.Cache) (PokemonLocation, error) {
	fullUrl := baseUrl + location
	reqCache, ok := c.Get(fullUrl)
	var pokemonLocation PokemonLocation

	if ok {
		err := json.Unmarshal(reqCache, &pokemonLocation)
		if err != nil {
			return pokemonLocation, err
		}

		return pokemonLocation, nil
	}


	res, err := http.Get(fullUrl)
	if err != nil {
		return pokemonLocation, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemonLocation, err
	}

	if err := json.Unmarshal(data, &pokemonLocation); err != nil {
		return pokemonLocation, err
	}

	c.Add(fullUrl, data)

	return pokemonLocation, nil
}


func getPokemon(pokemon, baseUrl string) (Pokemon, error) {
	fullUrl := baseUrl + pokemon
	var poke Pokemon

	res, err := http.Get(fullUrl)
	if err != nil {
		return poke, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return poke, err
	}

	if err := json.Unmarshal(data, &poke); err != nil {
		return poke, err
	}

	return poke, nil
}