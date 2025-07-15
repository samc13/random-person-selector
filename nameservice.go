package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
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

func GetPeople() People {
	filePath := "names.csv"
	return constructPeople(filePath)
}

func constructPeople(filePath string) People {
	records := readCsvFile(filePath)
	var persons []Person
	for _, record := range records[1:] { // Skip header
		person := parseToPerson(record)
		persons = append(persons, person)
	}
	return People{people: persons}
}

func readCsvFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read csv: %s", err)
	}

	return records
}

func parseToPerson(record []string) Person {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		log.Fatalf("failed to parse ID: %s", err)
	}
	addedOn, err := time.Parse("2006-01-02", record[2])
	if err != nil {
		log.Fatalf("failed to parse AddedOn date: %s", err)
	}
	var voidedOn *time.Time
	if record[3] != "NULL" {
		v, _ := time.Parse("2006-01-02", record[3])
		voidedOn = &v
	}

	return Person{
		ID:       id,
		Name:     record[1],
		AddedOn:  addedOn,
		VoidedOn: voidedOn,
	}
}
