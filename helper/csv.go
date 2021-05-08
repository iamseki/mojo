package helper

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func CreateCSVFile(filename string, header []string) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln("Error on call CreateCSVFile:", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()
	w.Write(header)

	return nil
}

func NextCSVRecord(r *csv.Reader) []string {
	record, err := r.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Fatalln("Error on call NextCSVRecord:", err)
	}
	return record
}

func GetCSVRecords(r *csv.Reader, size int) [][]string {
	// throw away header
	r.Read()
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Error on call GETCSVRecords:", err)
	}
	return records
}
