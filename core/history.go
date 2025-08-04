package core

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

func GetSelectedIDsThisYear(filename string, year int) (map[int]struct{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	selected := make(map[int]struct{})
	for _, rec := range records[1:] { // skip header
		if len(rec) < 2 {
			continue
		}
		id, err := strconv.Atoi(rec[0])
		if err != nil {
			continue
		}
		date, err := time.Parse("2006-01-02", rec[1])
		if err != nil {
			continue
		}
		if date.Year() == year {
			selected[id] = struct{}{}
		}
	}
	return selected, nil
}
