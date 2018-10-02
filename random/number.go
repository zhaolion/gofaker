package random

// Number will generate a random number between given min And max
func Number(min int, max int) int {
	return rangeInt(min, max)
}

// Numerify will replace # with random numerical values
func Numerify(str string) string {
	return replaceNumberWithPound(str)
}
