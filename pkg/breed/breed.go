package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Breed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getBreeds() ([]Breed, error) {
	url := "https://api.thecatapi.com/v1/breeds"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var breeds []Breed
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds, nil
}

func ValidateBreedName(breedName string) (bool, error) {
	breeds, err := getBreeds()
	if err != nil {
		return false, err
	}

	for _, breed := range breeds {
		if strings.ToLower(breed.Name) == strings.ToLower(breedName) {
			return true, nil
		}
	}
	return false, nil
}
