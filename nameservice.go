package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"errors"
	"time"
)

type Person struct {
	ID       int
	Name     string
	AddedOn  time.Time
	VoidedOn *time.Time // Nullable
}

func (p Person) String() string {
	if (p.VoidedOn != nil) {
		return p.Name + " (Voided)"
	}
	return p.Name
}

type People struct {
	people []Person
}

func GetPeople() (People, error) {
	filePath := "names.csv"
	return constructPeople(filePath)
}

func constructPeople(filePath string) (People, error) {
	records, err := readCsvFile(filePath)
	if err != nil {
		return People{}, err
	}
	if len(records) < 2 {
		return People{}, errors.New("CSV file is empty or does not contain enough data")
	}
	var persons []Person
	for _, record := range records[1:] { // Skip header
		person, err := parseToPerson(record)
		if err != nil {
			return People{}, err
		}
		persons = append(persons, person)
	}
	return People{people: persons}, nil
}

func readCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func parseToPerson(record []string) (Person, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return Person{}, err
	}
	addedOn, err := time.Parse("2006-01-02", record[2])
	if err != nil {
		return Person{}, err
	}
	var voidedOn *time.Time
	if record[3] != "NULL" {
		v, err := time.Parse("2006-01-02", record[3])
		if err != nil {
			return Person{}, err
		} else {
			voidedOn = &v
		}
	}

	return Person{
		ID:       id,
		Name:     record[1],
		AddedOn:  addedOn,
		VoidedOn: voidedOn,
	}, nil
}
