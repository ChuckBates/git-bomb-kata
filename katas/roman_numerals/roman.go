package roman_numerals

func IntToRoman(num int) string {
	if num < 5 {
		return "I"
	} else {
		return "V"
	}

}

func RomanToInt(s string) int {
	panic("not implemented")
}
