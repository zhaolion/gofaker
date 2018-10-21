package gofaker

import (
	"fmt"

	"github.com/zhaolion/gofaker/random"
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
			expect: "15116845420",
		},
		{
			locale: "bg",
			expect: "0871684542",
		},
		{
			locale: "ca-CAT",
			expect: "616.845.420",
		},
		{
			locale: "ca",
			expect: "(168) 454-2055",
		},
		{
			locale: "da-DK",
			expect: "30 16 84 54",
		},
		{
			locale: "de-AT",
			expect: "0616-84542055",
		},
		{
			locale: "de-CH",
			expect: "(168) 454-2055",
		},
		{
			locale: "de",
			expect: "+49-161-6845420",
		},
		{
			locale: "ee",
			expect: "516 8454",
		},
		{
			locale: "en-au-ocker",
			expect: "0416-845-420",
		},
		{
			locale: "en-AU",
			expect: "0416-845-420",
		},
		{
			locale: "en-BORK",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-GB",
			expect: "07516 845420",
		},
		{
			locale: "en-IND",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-MS",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-NEP",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-NG",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-NZ",
			expect: "0216 845 420",
		},
		{
			locale: "en-PAK",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-SG",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-UG",
			expect: "256 781 684 542",
		},
		{
			locale: "en-ZA",
			expect: "072 168 4542",
		},
		{
			locale: "es",
			expect: "616.845.420",
		},
		{
			locale: "fa",
			expect: "(168) 454-2055",
		},
		{
			locale: "fi-FI",
			expect: "041-6845420",
		},
		{
			locale: "fr-CH",
			expect: "+41.1684.542055",
		},
		{
			locale: "fr",
			expect: "07 16 84 54 20",
		},
		{
			locale: "he",
			expect: "041-684-5420",
		},
		{
			locale: "id",
			expect: "(168) 454-2055",
		},
		{
			locale: "it",
			expect: "(168) 454-2055",
		},
		{
			locale: "ja",
			expect: "080-1684-5420",
		},
		{
			locale: "ko",
			expect: "(168) 454-2055",
		},
		{
			locale: "lv",
			expect: "241 684 542",
		},
		{
			locale: "nb-NO",
			expect: "16 84 54 20",
		},
		{
			locale: "nl",
			expect: "(168) 454-2055",
		},
		{
			locale: "no",
			expect: "(168) 454-2055",
		},
		{
			locale: "pl",
			expect: "51-168-45-42",
		},
		{
			locale: "pt",
			expect: "(168) 454-2055",
		},
		{
			locale: "ru",
			expect: "(168) 454-2055",
		},
		{
			locale: "sk",
			expect: "(168) 454-2055",
		},
		{
			locale: "sv",
			expect: "076-168-4542",
		},
		{
			locale: "tr",
			expect: "(168) 454-2055",
		},
		{
			locale: "uk",
			expect: "(097) 168-45-42",
		},
		{
			locale: "vi",
			expect: "0167 168 4542",
		},
		{
			locale: "zh-TW",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-CA",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-US",
			expect: "(226) 614-8813",
		},
		{
			locale: "es-MX",
			expect: "433 684 5420",
		},
		{
			locale: "fr-CA",
			expect: "(168) 454-2055",
		},
		{
			locale: "pt-BR",
			expect: "(64) 91088-1364",
		},
	}

	for _, test := range tests {
		random.Seed(1000)
		bundle.SetLocale(test.locale)
		suite.Equal(test.expect, RandomCellPhone(), fmt.Sprintf("case: %s", test.locale))
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
			expect: "1684-54205516",
		},
		{
			locale: "bg",
			expect: "(168) 454-2055",
		},
		{
			locale: "ca-CAT",
			expect: "916.845.420",
		},
		{
			locale: "ca",
			expect: "(168) 454-2055",
		},
		{
			locale: "da-DK",
			expect: "16-84-54-20",
		},
		{
			locale: "de-AT",
			expect: "0168454205",
		},
		{
			locale: "de-CH",
			expect: "+41 16 845 42 05",
		},
		{
			locale: "de",
			expect: "(01684) 5420551",
		},
		{
			locale: "ee",
			expect: "661 6845",
		},
		{
			locale: "en-au-ocker",
			expect: "+61 1 6845 4205",
		},
		{
			locale: "en-AU",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-BORK",
			expect: "(168) 454-2055",
		},
		{
			locale: "en-GB",
			expect: "0916 845 4205",
		},
		{
			locale: "en-IND",
			expect: "+911684542055",
		},
		{
			locale: "en-MS",
			expect: "+60116845420",
		},
		{
			locale: "en-NEP",
			expect: "+977-1-6845420",
		},
		{
			locale: "en-NG",
			expect: "07016845420",
		},
		{
			locale: "en-NZ",
			expect: "+64 1 684 5420",
		},
		{
			locale: "en-PAK",
			expect: "+92 168 4542055",
		},
		{
			locale: "en-SG",
			expect: "+65 9168 4542",
		},
		{
			locale: "en-UG",
			expect: "+256 391 684 542",
		},
		{
			locale: "en-ZA",
			expect: "+2749 168 4542",
		},
		{
			locale: "es",
			expect: "916.845.420",
		},
		{
			locale: "fa",
			expect: "(168) 454-2055",
		},
		{
			locale: "fi-FI",
			expect: "168-4542055",
		},
		{
			locale: "fr-CH",
			expect: "+41 4 16 84 54 20",
		},
		{
			locale: "fr",
			expect: "02 16 84 54 20",
		},
		{
			locale: "he",
			expect: "04-168-4542",
		},
		{
			locale: "id",
			expect: "0816-8454-2055",
		},
		{
			locale: "it",
			expect: "316 845 420",
		},
		{
			locale: "ja",
			expect: "0168-45-4205",
		},
		{
			locale: "ko",
			expect: "016-845-4205",
		},
		{
			locale: "lv",
			expect: "61 684 542",
		},
		{
			locale: "nb-NO",
			expect: "16 84 54 20",
		},
		{
			locale: "nl",
			expect: "1684542055",
		},
		{
			locale: "no",
			expect: "(168) 454-2055",
		},
		{
			locale: "pl",
			expect: "82-168-45-42",
		},
		{
			locale: "pt",
			expect: "+351 (16) 845-4205",
		},
		{
			locale: "ru",
			expect: "+7(941)684-54-20",
		},
		{
			locale: "sk",
			expect: "016 8454 2055",
		},
		{
			locale: "sv",
			expect: "1684-542055",
		},
		{
			locale: "tr",
			expect: "90.168.454.2055",
		},
		{
			locale: "uk",
			expect: "(097) 168-45-42",
		},
		{
			locale: "vi",
			expect: "0916 845 4205",
		},
		{
			locale: "zh-TW",
			expect: "1684-54205516",
		},
		{
			locale: "en-CA",
			expect: "(579) 614-8813",
		},
		{
			locale: "en-US",
			expect: "(226) 614-8813",
		},
		{
			locale: "es-MX",
			expect: "433 684 5420",
		},
		{
			locale: "fr-CA",
			expect: "(579) 614-8813",
		},
		{
			locale: "pt-BR",
			expect: "(64) 1088-1364",
		},
	}

	for i, test := range tests {
		random.Seed(1000)
		bundle.SetLocale(test.locale)
		suite.Equal(test.expect, RandomPhoneNumber(), fmt.Sprintf("case: %d", i))
	}
}

func (suite *Suite) TestPhone() {
	bundle.SetLocale("en-US")
	suite.Equal("276", Phone().AreaCode())
	suite.Equal("229", Phone().ExchangeCode())
	suite.Equal("0881", Phone().SubscriberNumber())
	suite.Equal("3645", Phone().Extension())
}
