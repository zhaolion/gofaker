package random

import "math/rand"

// FixedBoolean random boolean with fixed ratio
// ratio should in [0.0, 1.0]
// - if ratio <= 0.0 always return false
// - if ratio >= 1.0 always return true
func FixedBoolean(ratio float32) bool {
	if ratio == 1.0 {
		return true
	} else if ratio == 0.0 {
		return false
	}

	return rand.Float32() < ratio
}
