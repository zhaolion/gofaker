package random

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixedBoolean(t *testing.T) {
	var tests = []struct {
		ratio  float32
		expect bool
	}{
		{
			ratio:  0.5,
			expect: false,
		},
		{
			ratio:  0.2,
			expect: false,
		},
		{
			ratio:  0.78,
			expect: true,
		},
		{
			ratio:  0.91,
			expect: true,
		},
		{
			ratio:  0.99,
			expect: true,
		},
		{
			ratio:  0.0,
			expect: false,
		},
		{
			ratio:  -0.1,
			expect: false,
		},
		{
			ratio:  1.0,
			expect: true,
		},
		{
			ratio:  1.2,
			expect: true,
		},
	}

	for i, test := range tests {
		Seed(1000)
		assert.Equal(t, test.expect, FixedBoolean(test.ratio), fmt.Sprintf("case: %d", i))
	}
}
