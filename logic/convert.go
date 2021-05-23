package logic

import (
	"encoding/csv"
	"mojo/contract"
	"mojo/helper"
	"os"
)

type ConvertOptions struct {
	from     string
	to       string
	contract string
	file     string
}

type Converter interface {
	Convert() error
}

type fileConverter struct {
	options ConvertOptions
}

func NewConvertLogic(options ConvertOptions) Converter {
	return &fileConverter{options}
}

func (fc fileConverter) Convert() error {
	err := fc.executeFromToConvertion()
	if err != nil {
		return err
	}
	return nil
}

func (fc fileConverter) executeFromToConvertion() error {
	convertFunction := fc.getFromToFunction()

	err := convertFunction()
	if err != nil {
		return err
	}

	return nil
}

func (fc fileConverter) getFromToFunction() func() error {
	fromToFunctions := map[string]func() error{
		"fromjsontocsv": fc.fromJsonToCsv,
		"fromcsvtojson": fc.fromCsvToJson,
	}

	return fromToFunctions["from"+fc.options.from+"to"+fc.options.to]
}

func (fc fileConverter) getContractToConvert() interface{} {
	switch fc.options.contract {
	case "b3":
		var b3 []contract.B3Price
		return b3
	default:
		return nil
	}
}

func (fc fileConverter) getCSVInputFromContractData(contractData interface{}) (rows [][]string, header []string) {
	// TO DO
	// use reflect package to analyze each position of contractData slice
	// and get its value
	// this tutorial might be helpful: https://golangbot.com/reflection/
	return rows, header
}

func (fc fileConverter) fromJsonToCsv() error {
	contractData := fc.getContractToConvert()
	jsonFileName := fc.options.file + ".json"
	csvFileName := fc.options.file + ".csv"

	err := helper.JSONFileToStruct(jsonFileName, &contractData)
	if err != nil {
		return err
	}

	rows, header := fc.getCSVInputFromContractData(contractData)

	err = helper.CreateCSVFile(csvFileName, header)
	if err != nil {
		return err
	}

	file, _ := os.OpenFile(csvFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer file.Close()

	w := csv.NewWriter(file)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}
	return nil
}

func (fc fileConverter) fromCsvToJson() error {
	return nil
}
