package roman_numerals

func IntToRoman(num int) string {
	if num < 5 {
		str := ""
		for range num {
			str += "I"
		}
		return str
	} else {
		return "V"
	}

}

func RomanToInt(s string) int {
	panic("not implemented")
}
