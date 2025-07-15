package main

import (
	"fmt"
	"log"
    "math/rand"
    "sort"
    "time"
)

func main() {
	people := GetPeople()

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

	// Take the top 10, or all if less than 10
	if len(peopleToUse) > 10 {
		peopleToUse = peopleToUse[:10]
	}

	log.Printf("Using %v people for selection", peopleToUse)

	// Randomly pick from the remaining sample
    rand.Seed(time.Now().UnixNano()) // Add random seed
	randomIndex := rand.Intn(len(peopleToUse))
	selectedPerson := peopleToUse[randomIndex]

	fmt.Printf("Selected person: %s \n", selectedPerson.Name)
}
