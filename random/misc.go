package random

import (
	"math/rand"
)

// rangeInt Generate random integer between min and max
func rangeInt(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn((max+1)-min) + min
}

// replaceNumberWithPound Replace # with number
func replaceNumberWithPound(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == poundTag {
			bytestr[i] = numberStringBytes[rand.Intn(9)]
		}
	}

	return string(bytestr)
}

// replaceLetterWithQuestionMark Replace ? with letter
func replaceLetterWithQuestionMark(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == quetionMarkTag {
			bytestr[i] = letterStringBytes[rand.Intn(26)]
		}
	}

	return string(bytestr)
}

func fixedString(l uint, source string) string {
	s := make([]byte, l)
	for i := 0; i < int(l); i++ {
		s[i] = source[rand.Intn(len(source))]
	}
	return string(s)
}
