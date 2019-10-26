package swapiGraphQLGo

import (
	"encoding/json"
)

type AllPeople struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []*Person   `json:"results"`
}

type AllPlanets struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []*Planet   `json:"results"`
}

// GetPlanet docs
func GetPlanet(id string) (*Planet, error) {
	planet := &Planet{}

	url := ConstructURL("planets", id)
	body, err := HandleRequest(url)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(body), &planet)

	return planet, nil
}

// GetPerson doc
func GetPerson(id string) (*Person, error) {
	person := &Person{}
	home := &Homeworld{}
	url := ConstructURL("people", id)
	body, err := HandleRequest(url)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(body), &person)
	json.Unmarshal([]byte(body), &home)

	planet, err := GetPlanet(SplitLinkID(home.Homeworld, "planets/"))
	person.Homeworld = planet

	return person, nil

}

// GetPeople doc
func GetPeople() ([]*Person, error) {
	AllPeople := &AllPeople{}
	people := []*Person{}
	url := ConstructURL("people", "")
	body, err := HandleRequest(url)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(body), &AllPeople)

	for _, character := range AllPeople.Results {
		people = append(people, character)
	}

	return people, nil
}

// GetPlanets doc
func GetPlanets() ([]*Planet, error) {
	allPlanets := &AllPlanets{}
	planet := []*Planet{}
	url := ConstructURL("planets", "")
	body, err := HandleRequest(url)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(body), &allPlanets)

	for _, value := range allPlanets.Results {
		planet = append(planet, value)
	}

	return planet, nil
}
