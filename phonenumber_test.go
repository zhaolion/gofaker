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
		{
			locale: "bg",
			expect: "1870467117",
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
			expect: "487.227.3851",
		},
		{
			locale: "bg",
			expect: "312.342.3336",
		},
		{
			locale: "zh-CN",
			expect: "24654634761",
		},
	}

	for i, test := range tests {
		bundle.SetLocale(test.locale)
		assert.Equal(t, test.expect, RandomPhoneNumber(), fmt.Sprintf("case: %d", i))
	}
}
