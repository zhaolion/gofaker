package faker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomCellPhone(t *testing.T) {
	var tests = []struct {
		locale string
		expect string
	}{
		{
			locale: "en",
			expect: "(168) 454-2055",
		},
		{
			locale: "zh-CN",
			expect: "17168514670",
		},
	}

	for i, test := range tests {
		bundle.SetLocale(test.locale)
		assert.Equal(t, test.expect, RandomCellPhone(), fmt.Sprintf("case: %d", i))
	}
}

func TestRandomPhoneNumber(t *testing.T) {
	var tests = []struct {
		locale string
		expect string
	}{
		{
			locale: "en",
			expect: "346-711-7084",
		},
		{
			locale: "zh-CN",
			expect: "2273-85170123",
		},
	}

	for i, test := range tests {
		bundle.SetLocale(test.locale)
		assert.Equal(t, test.expect, RandomPhoneNumber(), fmt.Sprintf("case: %d", i))
	}
}
