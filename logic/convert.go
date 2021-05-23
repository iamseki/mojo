package logic

import (
	"encoding/csv"
	"fmt"
	"mojo/contract"
	"mojo/helper"
	"os"
	"reflect"
)

type ConvertOptions struct {
	From     string
	To       string
	Contract string
	File     string
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

	return fromToFunctions["from"+fc.options.From+"to"+fc.options.To]
}

func (fc fileConverter) getContractToConvert() interface{} {
	switch fc.options.Contract {
	case "b3":
		var b3 []contract.B3Price
		return b3
	default:
		return nil
	}
}

func (fc fileConverter) getCSVInputFromContractData(contractData interface{}) (rows [][]string, header []string) {
	contractType := reflect.ValueOf(contractData).Index(0).Interface()
	header = helper.GetFieldsNameFromMap(contractType)
	fmt.Println(header)
	rows = helper.GetValuesFromSliceMap(contractData)
	fmt.Println(rows)

	return rows, header
}

func (fc fileConverter) fromJsonToCsv() error {
	contractData := fc.getContractToConvert()
	jsonFileName := fc.options.File + ".json"
	csvFileName := fc.options.File + ".csv"

	err := helper.JSONFileToStruct(jsonFileName, &contractData)
	if err != nil {
		return err
	}

	fmt.Println(contractData)

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
