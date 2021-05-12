package helper_test

import (
	"mojo/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceFromAWSUnformattedString(t *testing.T) {
	unformatted := `
	PRE ffa2f82c-3de9-4a61-b7b4-dd471e62fda6/
	PRE baa2f82c-3de9-4a61-b7b4-dd471e62fs54/
	PRE cfa2f82c-3de9-4a61-a7b4-dd471e62fda6/
	`

	formatted := helper.SliceFromAWSUnformattedString(unformatted)

	assert.Equal(t, "ffa2f82c-3de9-4a61-b7b4-dd471e62fda6", formatted[0])
	assert.Equal(t, "baa2f82c-3de9-4a61-b7b4-dd471e62fs54", formatted[1])
	assert.Equal(t, "cfa2f82c-3de9-4a61-a7b4-dd471e62fda6", formatted[2])
}
