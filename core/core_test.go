package core

import (
	"strconv"
	"testing"
	"time"
)

type MockPeopleProvider struct {
	people People
	err    error
}

func (m MockPeopleProvider) GetPeople() (People, error) {
	return m.people, m.err
}

func givenTestPeople() []Person {
	return []Person{
		givenTestPerson(1),
		givenTestPerson(2),
		givenTestPerson(3),
		givenTestPerson(4),
		givenVoidedTestPerson(5),
	}
}

func givenTestPerson(i int) Person {
	return Person{
		ID:       i,
		Name:     "Test Person " + strconv.Itoa(i),
		AddedOn:  time.Now().Add(-time.Duration(2*i) * 24 * time.Hour),
		VoidedOn: nil,
	}
}

func givenVoidedTestPerson(i int) Person {
	return Person{
		ID:       i,
		Name:     "Test Person " + strconv.Itoa(i),
		AddedOn:  time.Now().Add(-time.Duration(2*i) * 24 * time.Hour),
		VoidedOn: &[]time.Time{time.Now().Add(-time.Duration(i) * 24 * time.Hour)}[0], // Set to current time to indicate voided
	}
}

func TestSelectRandomPerson_Success(t *testing.T) {
	// Arrange
	people := givenTestPeople()
	mockProvider := MockPeopleProvider{
		people: People{people: people},
		err:    nil,
	}

	// Act
	selectedPerson, err := FetchRandomPerson(mockProvider)

	// Assert
	assertNoError(&err, t)
	if selectedPerson == (Person{}) {
		t.Fatal("Expected a person to be selected, got an empty person")
	}
	if selectedPerson.VoidedOn != nil {
		t.Fatalf("Expected selected person to not be voided, got %v", selectedPerson.VoidedOn)
	}
}

func TestSelectRandomPerson_NoPeopleReturnedFromCsv(t *testing.T) {
	// Arrange
	mockProvider := MockPeopleProvider{
		people: People{people: []Person{}},
		err:    nil,
	}

	// Act
	selectedPerson, err := FetchRandomPerson(mockProvider)

	// Assert
	if err == nil {
		t.Fatal("Expected an error when no active people are available")
	}
	if selectedPerson != (Person{}) {
		t.Fatalf("Expected no person to be selected, got %v", selectedPerson)
	}
}

func TestSelectRandomPerson_EveryoneIsVoided(t *testing.T) {
	// Arrange
	people := []Person{
		givenVoidedTestPerson(1),
		givenVoidedTestPerson(2),
		givenVoidedTestPerson(3),
		givenVoidedTestPerson(4),
	}
	mockProvider := MockPeopleProvider{
		people: People{people: people},
		err:    nil,
	}

	// Act
	selectedPerson, err := FetchRandomPerson(mockProvider)

	// Assert
	if err == nil {
		t.Fatal("Expected an error when no active people are available")
	}
	if selectedPerson != (Person{}) {
		t.Fatalf("Expected no person to be selected, got %v", selectedPerson)
	}
}

func assertNoError(err *error, t *testing.T) {
	if *err != nil {
		t.Fatalf("%s", "Expected no error, got "+(*err).Error())
	}
}
