package helper_test

import (
	"encoding/csv"
	"mojo/helper"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testingFile = "testing.csv"

func TestCreateCSVFile(t *testing.T) {
	header := []string{"column1", "column2"}
	helper.CreateCSVFile(testingFile, header)

	file, err := os.OpenFile(testingFile, os.O_RDWR, 0644)

	assert.Nil(t, err)
	assert.Equal(t, testingFile, file.Name())

	csv := csv.NewReader(file)
	row, err := csv.Read()
	assert.Nil(t, err)
	assert.Equal(t, header, row)

	file.Close()

	err = os.Remove(testingFile)
	assert.Nil(t, err)
}

func TestNextCSVRecord(t *testing.T) {
	header := []string{"column1", "column2"}
	helper.CreateCSVFile(testingFile, header)

	file, _ := os.OpenFile(testingFile, os.O_RDWR, 0644)
	csv := csv.NewReader(file)

	record := helper.NextCSVRecord(csv)
	nextRecord := helper.NextCSVRecord(csv)

	var expectedNextRecord []string

	assert.Equal(t, header, record)
	assert.Equal(t, expectedNextRecord, nextRecord)
}

func TestGetCSVRecords(t *testing.T) {
	header := []string{"column1", "column2"}
	helper.CreateCSVFile(testingFile, header)

	file, _ := os.OpenFile(testingFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	data := [][]string{{"data1", "data2"}, {"data3", "data4"}}
	writer := csv.NewWriter(file)
	writer.WriteAll(data)

	file.Close()

	file, _ = os.OpenFile(testingFile, os.O_RDWR, 0644)
	reader := csv.NewReader(file)
	records := helper.GetCSVRecords(reader, 10)

	assert.Equal(t, data, records)

	os.Remove(testingFile)
}
