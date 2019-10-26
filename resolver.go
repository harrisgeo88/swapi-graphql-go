package swapiGraphQLGo

import (
	"context"
	"strconv"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) People(ctx context.Context) ([]*Person, error) {
	people, err := GetPeople()
	if err != nil {
		return nil, err
	}
	return people, nil
}

func (r *queryResolver) Planets(ctx context.Context) ([]*Planet, error) {
	planets, err := GetPlanets()
	if err != nil {
		return nil, err
	}
	return planets, nil
}

func (r *queryResolver) Planet(ctx context.Context, input int) (*Planet, error) {
	planet, err := GetPlanet(strconv.Itoa(input))
	if err != nil {
		return nil, err
	}
	return planet, nil
}

func (r *queryResolver) Person(ctx context.Context, input int) (*Person, error) {
	person, err := GetPerson(strconv.Itoa(input))
	if err != nil {
		return nil, err
	}
	return person, nil
}
