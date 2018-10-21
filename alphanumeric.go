package gofaker

import "github.com/zhaolion/gofaker/random"

// Alphanumeric string(mixed alpha with number) of fixed length
func Alphanumeric(l uint) string {
	return random.FixedAlphaNumeric(l)
}

// Alpha alpha string of fixed length
func Alpha(l uint) string {
	return random.FixedAlpha(l)
}

// Numberic numeric string of fixed length
func Numberic(l uint) string {
	return random.FixedNumeric(l)
}
