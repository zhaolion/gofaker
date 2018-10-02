package faker

import (
	"fmt"
)

func (suite *Suite) TestRandomCellPhone() {
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
		{
			locale: "ca-CAT",
			expect: "648722738",
		},
		{
			locale: "ca",
			expect: "(170) 123-4233",
		},
		{
			locale: "da-DK",
			expect: "20 66 22 46",
		},
		{
			locale: "de-AT",
			expect: "+43-646-3476133",
		},
		{
			locale: "de-CH",
			expect: "683-280-3530",
		},
		{
			locale: "de",
			expect: "+49-170-3655675",
		},
		{
			locale: "ee",
			expect: "562 0604",
		},
	}

	for i, test := range tests {
		bundle.SetLocale(test.locale)
		suite.Equal(test.expect, RandomCellPhone(), fmt.Sprintf("case: %d", i))
	}
}

func (suite *Suite) TestRandomPhoneNumber() {
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
			expect: "6851-46701046",
		},
		{
			locale: "bg",
			expect: "(117) 084-8722",
		},
		{
			locale: "ca-CAT",
			expect: "938.517.012",
		},
		{
			locale: "ca",
			expect: "423-336-6224",
		},
		{
			locale: "da-DK",
			expect: "54-63-47-61",
		},
		{
			locale: "de-AT",
			expect: "413568328",
		},
		{
			locale: "de-CH",
			expect: "153 020 36 55",
		},
		{
			locale: "de",
			expect: "(05162) 0604384",
		},
		{
			locale: "ee",
			expect: "758 7823",
		},
	}

	for i, test := range tests {
		bundle.SetLocale(test.locale)
		suite.Equal(test.expect, RandomPhoneNumber(), fmt.Sprintf("case: %d", i))
	}
}
