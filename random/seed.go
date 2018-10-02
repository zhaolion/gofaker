package random

import (
	"math/rand"
	"time"
)

// Seed random.
// if you set the same seed, you will get the same random serial sequence
func Seed(seed int64) {
	if seed == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else {
		rand.Seed(seed)
	}
}
