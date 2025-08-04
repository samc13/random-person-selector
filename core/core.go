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

	peopleSelectedThisYear, err := GetSelectedIDsThisYear("previous_selections.csv", time.Now().Year())
	if err != nil {
		log.Fatalf("Error getting selected IDs this year: %v", err)
	}
	filteredPeople, err := excludeSelectedPeople(peopleToUse, peopleSelectedThisYear)
	if err != nil {
		log.Fatalf("Error excluding selected people: %v", err)
	}

	sample, err := buildWeightedPool(filteredPeople)
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

func excludeSelectedPeople(people []Person, excluded map[int]struct{}) ([]Person, error) {
	var result []Person
	for _, p := range people {
		if _, found := excluded[p.ID]; !found {
			result = append(result, p)
		}
	}
	return result, nil
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
