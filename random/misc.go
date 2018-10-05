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

// replaceNumberWithPound Replace # with numbers
func replaceNumberWithPound(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	hashtag := []byte("#")[0]
	numbers := []byte("0123456789")
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == hashtag {
			bytestr[i] = numbers[rand.Intn(9)]
		}
	}
	if bytestr[0] == []byte("0")[0] {
		bytestr[0] = numbers[rand.Intn(8)+1]
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
