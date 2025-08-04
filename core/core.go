package core

import (
	"log"
	"sort"
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

	sortPeopleToUse, err := orderByOldestFirst(peopleToUse)
	if err != nil {
		log.Fatalf("Error ordering people: %v", err)
	}

	// Take the top maxSampleSize, or all if less than maxSampleSize
	sample, err := takeAtMostTen(sortPeopleToUse)
	if err != nil {
		log.Fatalf("Error taking sample: %v", err)
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

func orderByOldestFirst(people []Person) ([]Person, error) {
	copyOfPeople := make([]Person, len(people))
	copy(copyOfPeople, people)
	sort.Slice(copyOfPeople, func(i, j int) bool {
		return copyOfPeople[i].AddedOn.Before(copyOfPeople[j].AddedOn)
	})
	return copyOfPeople, nil
}

func takeAtMostTen(people []Person) ([]Person, error) {
	maxSampleSize := 10
	if len(people) > maxSampleSize {
		people = people[:maxSampleSize]
	}
	return people, nil
}
