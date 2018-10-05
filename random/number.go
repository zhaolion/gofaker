package random

import "math/rand"

// Number will generate a random number between given min And max
func Number(min int, max int) int {
	return rangeInt(min, max)
}

// Numerify will replace # with random numerical values
func Numerify(str string) string {
	return replaceNumberWithPound(str)
}

const chars = "0123456789"

// FixedNumberString return number string of fixed length
func FixedNumberString(l uint) string {
	s := make([]byte, l)
	for i := 0; i < int(l); i++ {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}
