package helper

import (
	"encoding/csv"
	"os"
)

func CreateCSVFile(filename string, header []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()
	w.Write(header)

	return nil
}

func NextCSVRecord(r *csv.Reader) ([]string, error) {
	record, err := r.Read()
	if err != nil {
		return nil, err
	}
	return record, nil
}

func CSVRecords(r *csv.Reader, size int) ([][]string, error) {
	// throw away header
	r.Read()
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
