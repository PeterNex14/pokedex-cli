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



func  getLocationArea(baseUrl string, c *pokecache.Cache) (ListLocation, error) {
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