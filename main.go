package main

import (
	"log"
	"random-person-selector/core"
	"random-person-selector/formatting"
	"random-person-selector/logging"
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
	logging.Infof("Selected person: %s", formatting.FormatName(selectedPerson.Name))

	logging.Debug("Storing selection...")
	if err := core.RecordSelection(selectedPerson, "previous_selections.csv"); err != nil {
		log.Fatalf("Error storing selection: %v", err)
	}
	logging.Debug("Selection stored successfully.")
}
