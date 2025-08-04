package core

import (
	"fmt"
	"math/rand"
	"time"
)

func SelectRandomPerson(people []Person) (Person, error) {
	if len(people) == 0 {
		return Person{}, fmt.Errorf("no people available for selection")
	}
	// Randomly pick from the valid people
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(people))
	selectedPerson := people[randomIndex]

	return selectedPerson, nil
}
