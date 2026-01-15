package main

import (
	"encoding/json"
	"net/http"
)

type ListLocation struct {
	Next		*string 	`json:"next"`
	Previous 	*string		`json:"previous"`
	Results		[]struct {
		Name	string	`json:"name"`
		URL		string	`json:"url"`
	}	`json:"results"`
}



func getLocationArea(baseUrl string) (ListLocation, error) {
	res, err := http.Get(baseUrl)

	if err != nil {
		return ListLocation{},  err
	}

	defer res.Body.Close()

	var listLocation ListLocation
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&listLocation); err != nil {
		return ListLocation{}, err
	}

	return listLocation, nil

}