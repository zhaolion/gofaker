package random

// FixedAlpha return alpha string of fixed length
// mixed with lower and higher case string
func FixedAlpha(l uint) string {
	return fixedString(l, alphaNumerics[:52])
}

// FixedNumeric return number string of fixed length
func FixedNumeric(l uint) string {
	return fixedString(l, alphaNumerics[52:])
}

// FixedAlphaNumeric return string(mixed alpha with number) of fixed length
func FixedAlphaNumeric(l uint) string {
	return fixedString(l, alphaNumerics)
}

// Numerify replace # with random numerical values
func Numerify(str string) string {
	return replaceNumberWithPound(str)
}

// Letterify replace ? with random letter
func Letterify(str string) string {
	return replaceLetterWithQuestionMark(str)
}

// Bothify replace both tags
// replace # with random numerical values
// replace ? with random letter
func Bothify(str string) string {
	return Letterify(Numerify(str))
}
