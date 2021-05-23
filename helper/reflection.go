package helper

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GetValuesFromSliceMap(any interface{}) [][]string {
	values := [][]string{}

	if reflect.ValueOf(any).Kind() == reflect.Slice {
		len := reflect.ValueOf(any).Len()
		for i := 0; i < len; i++ {
			// for each position in the slice
			obj := reflect.ValueOf(any).Index(i)
			fmt.Println(obj)
			m := obj.Elem()
			keys := m.MapKeys()

			valuesFromObj := []string{}
			for _, k := range keys {
				//v := m.MapIndex(k.Convert(m.Type().Key()))

				switch v := m.MapIndex(k.Convert(m.Type().Key())).Interface().(type) {
				case string:
					// convert to string
					valuesFromObj = append(valuesFromObj, v)
				case int:
					valuesFromObj = append(valuesFromObj, strconv.Itoa(v))
				case bool:
					valuesFromObj = append(valuesFromObj, strconv.FormatBool(v))
				case time.Time:
					valuesFromObj = append(valuesFromObj, v.Format("2006-01-02"))
				}
			}

			values = append(values, valuesFromObj)
		}
	}

	return values
}

func GetFieldsNameFromMap(any interface{}) []string {
	fields := []string{}
	// get type of interface{} at runtime
	switch reflect.ValueOf(any).Kind() {
	// if were a struct
	case reflect.Map:
		// get the struct value and how many fields the struct have
		v := reflect.ValueOf(any)
		keys := v.MapKeys()
		// for each field, get the name of the field
		for _, key := range keys {
			field := key.Convert(reflect.TypeOf(any).Key()).String()
			fields = append(fields, field)
		}
	}

	return fields
}

func GetFieldsNameFromStruct(any interface{}) []string {
	fields := []string{}
	// get type of interface{} at runtime
	switch reflect.ValueOf(any).Kind() {
	// if were a struct
	case reflect.Struct:
		// get the struct value and how many fields the struct have
		v := reflect.ValueOf(any)
		numFields := v.NumField()
		// for each field, get the name of the field
		for i := 0; i < numFields; i++ {
			fieldName := v.Type().Field(i).Name
			fields = append(fields, strings.ToLower(fieldName))
		}
	}

	return fields
}

func GetValuesFromSliceStruct(any interface{}) [][]string {
	values := [][]string{}

	if reflect.ValueOf(any).Kind() == reflect.Slice {
		len := reflect.ValueOf(any).Len()
		for i := 0; i < len; i++ {
			// for each position in the slice
			obj := reflect.ValueOf(any).Index(i)
			propertiesNum := obj.NumField()
			valuesFromObj := []string{}
			// get the values from object
			for k := 0; k < propertiesNum; k++ {
				switch v := obj.Field(k).Interface().(type) {
				case string:
					// convert to string
					valuesFromObj = append(valuesFromObj, v)
				case int:
					valuesFromObj = append(valuesFromObj, strconv.Itoa(v))
				case bool:
					valuesFromObj = append(valuesFromObj, strconv.FormatBool(v))
				case time.Time:
					valuesFromObj = append(valuesFromObj, v.Format("2006-01-02"))
				}
			}
			// add to slice containing al values
			values = append(values, valuesFromObj)
		}
	}

	return values
}
