package helper_test

import (
	"mojo/helper"
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

	var sut random

	fields := helper.GetFieldsNameFromStruct(sut)

	assert.Contains(t, fields, "id")
	assert.Contains(t, fields, "name")
	assert.Contains(t, fields, "age")
	assert.Contains(t, fields, "date")
}
