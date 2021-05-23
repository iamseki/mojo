package helper

import (
	"reflect"
	"strings"
)

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
