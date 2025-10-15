package roman_numerals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	input := 1
	expected := "I"

	actual := IntToRoman(input)

	assert.Equal(t, expected, actual)
}
