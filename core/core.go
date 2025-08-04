package core

import (
	"log"
	"time"
)

type PeopleProvider interface {
	GetPeople() (People, error)
}

func FetchRandomPerson(provider PeopleProvider) (Person, error) {
	people, err := provider.GetPeople()
	if err != nil {
		log.Fatalf("Error getting people: %v", err)
	}

	peopleToUse, err := filterOutVoided(people.people)
	if err != nil {
		log.Fatalf("Error filtering voided persons: %v", err)
	}

	// TODO: Filter down to those that have been selected the least (need to incorporate previousslections.csv)

	sample, err := buildWeightedPool(peopleToUse)
	if err != nil {
		log.Fatalf("Error ordering people: %v", err)
	}

	result, err := SelectRandomPerson(sample)

	return result, err
}

func filterOutVoided(people []Person) ([]Person, error) {
	var output []Person
	for _, person := range people {
		if person.VoidedOn == nil {
			output = append(output, person)
		}
	}
	return output, nil
}

// Returns a weighted pool of people based on years present
func buildWeightedPool(people []Person) ([]Person, error) {
	now := time.Now()
	var pool []Person
	for _, p := range people {
		years := int(now.Sub(p.AddedOn).Hours()/(24*365)) + 1 // at least 1
		// Create as many entries as years of service
		for i := 0; i < years; i++ {
			pool = append(pool, p)
		}
	}
	return pool, nil
}
