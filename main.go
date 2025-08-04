package main

import (
	"fmt"
	"log"
	"random-person-selector/core"
)

type FileBasedPeopleProvider struct{}

func (f FileBasedPeopleProvider) GetPeople() (core.People, error) {
	return core.GetPeople()
}

func main() {
	// Initialize the people provider
	provider := FileBasedPeopleProvider{}
	selectedPerson, err := core.FetchRandomPerson(provider)

	if err != nil {
		log.Fatalf("Error selecting random person: %v", err)
	}
	fmt.Printf("Selected person: %s \n", selectedPerson.Name)

	fmt.Printf("Storing selection...\n")
	if err := core.RecordSelection(selectedPerson, "previous_selections.csv"); err != nil {
		log.Fatalf("Error storing selection: %v", err)
	}
	fmt.Printf("Selection stored successfully.\n")
}
