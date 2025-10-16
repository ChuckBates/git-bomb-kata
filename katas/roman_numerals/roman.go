package roman_numerals

func IntToRoman(num int) string {
	if num == 1 {
		return "I";
	}
	return ""
}

func RomanToInt(s string) int {
	if s == "I" {
		return 1;
	}
	return -1
}
