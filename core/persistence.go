package core

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

func RecordSelection(person Person, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		strconv.Itoa(person.ID),
		time.Now().Format("2006-01-02"),
		"Added by random person selector",
	}
	return writer.Write(record)
}
