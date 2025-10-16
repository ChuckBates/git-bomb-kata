package roman_numerals

func IntToRoman(num int) string {
	if num < 5 {
		if num == 4 {
			return "IV"
		}

	} else if num < 10 {
		remainder := num % 5
		return "V"
	} else if num == 10 {
		return "X"
	} else {
		return "V"
	}
}

func getI(num int) string {
	str := ""
	for range num {
		str += "I"
	}
	return str
}

func RomanToInt(s string) int {
	panic("not implemented")
}
