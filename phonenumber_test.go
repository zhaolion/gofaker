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
		{
			locale: "en-PAK",
			expect: "741.474.1466",
		},
		{
			locale: "en-SG",
			expect: "1-565-457-5560",
		},
		{
			locale: "en-UG",
			expect: "7 747 872 553",
		},
		{
			locale: "en-ZA",
			expect: "681 226 0108",
		},
		{
			locale: "es",
			expect: "610873768",
		},
		{
			locale: "fa",
			expect: "1-638-645-2100",
		},
		{
			locale: "fi-FI",
			expect: "170-5777503",
		},
		{
			locale: "fr-CH",
			expect: "+41.8025.070354",
		},
		{
			locale: "fr",
			expect: "+33 613244425",
		},
		{
			locale: "he",
			expect: "845-383-6848",
		},
		{
			locale: "id",
			expect: "(211) 583-8065",
		},
		{
			locale: "it",
			expect: "327-307-4687",
		},
		{
			locale: "ja",
			expect: "680-6111-5572",
		},
		{
			locale: "ko",
			expect: "601-270-0711",
		},
		{
			locale: "iv",
			expect: "568-818-0426",
		},
		{
			locale: "nb-NO",
			expect: "40865345",
		},
		{
			locale: "nl",
			expect: "1-151-365-3148",
		},
		{
			locale: "no",
			expect: "233.154.4031",
		},
		{
			locale: "pl",
			expect: "79-587-00-78",
		},
		{
			locale: "pt",
			expect: "(880) 060-4267",
		},
		{
			locale: "ru",
			expect: "1-711-502-5463",
		},
		{
			locale: "sk",
			expect: "650.638.0072",
		},
		{
			locale: "sv",
			expect: "476-306-7140",
		},
		{
			locale: "tr",
			expect: "240-563-4276",
		},
		{
			locale: "uk",
			expect: "(092) 867-85-50",
		},
		{
			locale: "vi",
			expect: "8186 236 8615",
		},
		{
			locale: "zh-TW",
			expect: "(533) 317-8856",
		},
		{
			locale: "en-CA",
			expect: "(144) 722-2381",
		},
		{
			locale: "en-US",
			expect: "518.734.6775",
		},
		{
			locale: "en-MX",
			expect: "644 247 156 6020",
		},
		{
			locale: "fr-CA",
			expect: "(122) 041-6741",
		},
		{
			locale: "pt-BR",
			expect: "(92) 98576-5803",
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
		{
			locale: "en-PAK",
			expect: "4356-5457556",
		},
		{
			locale: "en-SG",
			expect: "+65 8787 2553",
		},
		{
			locale: "en-UG",
			expect: "+256 418 226 010",
		},
		{
			locale: "en-ZA",
			expect: "+2728 821 0873",
		},
		{
			locale: "es",
			expect: "968463864",
		},
		{
			locale: "fa",
			expect: "210.070.5777 x50360",
		},
		{
			locale: "fi-FI",
			expect: "825-0703543",
		},
		{
			locale: "fr-CH",
			expect: "51 24 44 25 45",
		},
		{
			locale: "fr",
			expect: "23 36 84 80 82",
		},
		{
			locale: "he",
			expect: "51-583-8065",
		},
		{
			locale: "id",
			expect: "282730746877",
		},
		{
			locale: "it",
			expect: "+11 557 28360127",
		},
		{
			locale: "ja",
			expect: "30-7114-5688",
		},
		{
			locale: "ko",
			expect: "204-2654-0865",
		},
		{
			locale: "iv",
			expect: "452-151-3653",
		},
		{
			locale: "nb-NO",
			expect: "842 33 154",
		},
		{
			locale: "nl",
			expect: "(0311) 587007",
		},
		{
			locale: "no",
			expect: "188-006-0426 x7271",
		},
		{
			locale: "pl",
			expect: "23-502-54-63",
		},
		{
			locale: "pt",
			expect: "(65) 063-8007",
		},
		{
			locale: "ru",
			expect: "+7(924)306-71-40",
		},
		{
			locale: "sk",
			expect: "+421 024 056 342",
		},
		{
			locale: "sv",
			expect: "6586-785503",
		},
		{
			locale: "tr",
			expect: "903686153453",
		},
		{
			locale: "uk",
			expect: "(050) 317-88-56",
		},
		{
			locale: "vi",
			expect: "2914 472 2238",
		},
		{
			locale: "zh-TW",
			expect: "2678-74431566",
		},
		{
			locale: "en-CA",
			expect: "1-780-510-5799",
		},
		{
			locale: "en-US",
			expect: "(234) 213-5835",
		},
		{
			locale: "en-MX",
			expect: "55-2076-3664",
		},
		{
			locale: "fr-CA",
			expect: "289-731-6994",
		},
		{
			locale: "pt-BR",
			expect: "(28) 4108-7406",
		},
	}

	for i, test := range tests {
		bundle.SetLocale(test.locale)
		suite.Equal(test.expect, RandomPhoneNumber(), fmt.Sprintf("case: %d", i))
	}
}
