package core

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"
)

type PeopleProvider interface {
	GetPeople() (People, error)
}

func SelectRandomPerson(provider PeopleProvider) (Person, error) {
	people, err := provider.GetPeople()
	if err != nil {
		log.Fatalf("Error getting people: %v", err)
	}
	// Filter out those that are voided
	var peopleToUse []Person
	for _, person := range people.people {
		if person.VoidedOn == nil {
			peopleToUse = append(peopleToUse, person)
		}
	}
	log.Printf("Full list of people looks like %v", peopleToUse)

	// TODO: Filter down to those that have been selected the least (need to incorporate previousslections.csv)

	// Order by the 'oldest' addedOn date
	sort.Slice(peopleToUse, func(i, j int) bool {
		return peopleToUse[i].AddedOn.Before(peopleToUse[j].AddedOn)
	})

	log.Printf("Ordered array looks like %v", peopleToUse)

	// Take the top maxSampleSize, or all if less than maxSampleSize
	maxSampleSize := 10
	if len(peopleToUse) > maxSampleSize {
		peopleToUse = peopleToUse[:maxSampleSize]
	}

	log.Printf("Using %v people for selection", peopleToUse)

	// Ensure there are people to select from
	if len(peopleToUse) == 0 {
		return Person{}, fmt.Errorf("failed to find any people to select from")
	}
	// Randomly pick from the remaining sample
	rand.Seed(time.Now().UnixNano()) // Add random seed
	randomIndex := rand.Intn(len(peopleToUse))
	selectedPerson := peopleToUse[randomIndex]

	return selectedPerson, nil
}
