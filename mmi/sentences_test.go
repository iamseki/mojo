package mmi_test

import (
	"mojo/mmi"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFunnySentence(t *testing.T) {
	expectedSentence := "Algorithm: Word used by programmers when they don’t want to explain what they did."
	sentence := mmi.GetFunnySentence(6)
	nilSentence := mmi.GetFunnySentence(77)

	assert.Equal(t, expectedSentence, sentence)
	assert.Empty(t, nilSentence)
}

func TestGetRandomSentence(t *testing.T) {
	sentence := mmi.GetRandomSentence()
	assert.NotEmpty(t, sentence)
}
