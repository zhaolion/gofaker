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
		{
			locale: "en-au-ocker",
			expect: "8484-358-782",
		},
		{
			locale: "en-AU",
			expect: "8432 166 341",
		},
		{
			locale: "en-BORK",
			expect: "(768) 663-1402",
		},
		{
			locale: "en-GB",
			expect: "57977 735273",
		},
		{
			locale: "en-IND",
			expect: "(244) 105-0722",
		},
		{
			locale: "en-MS",
			expect: "1-655-054-7053",
		},
		{
			locale: "en-NEP",
			expect: "665-820-0152",
		},
		{
			locale: "en-NG",
			expect: "1-822-531-2753",
		},
		{
			locale: "en-NZ",
			expect: "4258 106 766",
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
		{
			locale: "en-au-ocker",
			expect: "23 2166 3418",
		},
		{
			locale: "en-AU",
			expect: "686.631.4025",
		},
		{
			locale: "en-BORK",
			expect: "773-527-3382 x44105",
		},
		{
			locale: "en-GB",
			expect: "4117 228 6550",
		},
		{
			locale: "en-IND",
			expect: "+917053806582",
		},
		{
			locale: "en-MS",
			expect: "+60101523282",
		},
		{
			locale: "en-NEP",
			expect: "+97753127530",
		},
		{
			locale: "en-NG",
			expect: "49081067663",
		},
		{
			locale: "en-NZ",
			expect: "24 147 4146",
		},
	}

	for i, test := range tests {
		bundle.SetLocale(test.locale)
		suite.Equal(test.expect, RandomPhoneNumber(), fmt.Sprintf("case: %d", i))
	}
}
