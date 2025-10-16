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

func TestIntToRomanIII(t *testing.T) {
	input := 3
	expected := "III"

	actual := IntToRoman(input)

	assert.Equal(t, expected, actual)
}

func TestIntToRomanIV(t *testing.T) {
	input := 4
	expected := "IV"

	actual := IntToRoman(input)

	assert.Equal(t, expected, actual)
}

func TestIntToRomanV(t *testing.T) {
	input := 5
	expected := "V"

	actual := IntToRoman(input)

	assert.Equal(t, expected, actual)
}

func TestIntToRomanX(t *testing.T) {
	input := 10
	expected := "X"

	actual := IntToRoman(input)

	assert.Equal(t, expected, actual)
}

func TestIntToRomanVII(t *testing.T) {
	input := 7
	expected := "VII"

	actual := IntToRoman(input)

	assert.Equal(t, expected, actual)
}
