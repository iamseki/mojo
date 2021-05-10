package helper_test

import (
	"mojo/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonthByLetterInOptionTickr(t *testing.T) {
	sut := map[string]int{
		"VALEM370": 1,
		"VALEA370": 1,
		"LOGNB145": 2,
		"ENATN125": 2,
		"CNTOO230": 3,
		"ELETC280": 3,
		"PETRD150": 4,
		"PETRP150": 4,
		"ABEVQ96":  5,
		"MGLUE545": 5,
		"SUZBR39":  6,
		"BBASF40":  6,
		"CSNAG78":  7,
		"CSNAS78":  7,
		"USIMH107": 8,
		"USIMT107": 8,
		"PETRI150": 9,
		"PETRU150": 9,
		"VALEJ485": 10,
		"VALEV485": 10,
		"PCARK20":  11,
		"PCARW20":  11,
		"PETRL622": 12,
		"PETRX622": 12,
	}

	for key, value := range sut {
		month := helper.MonthByLetterInOptionTickr(rune(key[4]))
		assert.Equal(t, value, month)
	}
}
