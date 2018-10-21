package gofaker

import (
	"github.com/zhaolion/gofaker/random"
)

// Boolean random boolean
// ratio is 0.5
func Boolean() bool {
	return random.FixedBoolean(0.5)
}
