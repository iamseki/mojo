package helper_test

import (
	"fmt"
	"mojo/helper"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPropertiesNameFromStruct(t *testing.T) {
	type random struct {
		ID   string
		Name string
		Age  int
		Date time.Time
	}

	suts := []random{{ID: "test", Name: "test"}}

	fields := helper.GetFieldsNameFromStruct(reflect.ValueOf(suts).Index(0).Interface())

	assert.Contains(t, fields, "id")
	assert.Contains(t, fields, "name")
	assert.Contains(t, fields, "age")
	assert.Contains(t, fields, "date")
}

func TestGetValuesFromSliceStruct(t *testing.T) {
	type random struct {
		ID        string
		Name      string
		Age       int
		Manager   bool
		UpdatedAt time.Time
	}

	sut := []random{
		{ID: "01", Name: "Thamires", Age: 24, Manager: true},
		{ID: "02", Name: "Rafaella", Age: 8, UpdatedAt: time.Now()},
	}

	values := helper.GetValuesFromSliceStruct(sut)

	fmt.Println(values)
}
