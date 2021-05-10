package helper_test

import (
	"mojo/helper"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMojoBrainPath(t *testing.T) {
	path, err := helper.MojoBrainPath()
	assert.Nil(t, err)

	home := os.Getenv("HOME")
	assert.Equal(t, home+"/.mojo-brain.db", path)
}
