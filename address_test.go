package gofaker

import (
	"fmt"

	"github.com/zhaolion/gofaker/random"
)

func (suite *Suite) TestFakeAddress() {
	suite.NotNil(FakeAddress())
}

func (suite *Suite) TestAddress() {
	var tests = []struct {
		locale           string
		City             string
		StreetName       string
		StreetAddress    string
		SecondaryAddress string
		BuildingNumber   string
		Community        string
		Postcode         string
		ZipCode          string
		TimeZone         string
		StreetSuffix     string
		CitySuffix       string
		CityPrefix       string
		State            string
		StateAbbr        string
		Country          string
		CountryCode      string
		CountryLongCode  string
		FullAddress      string
	}{
		{
			locale:           "en",
			City:             "North Ermelinda",
			StreetName:       "Rodger Forges",
			StreetAddress:    "55168 Dave Burg",
			SecondaryAddress: "Suite 010",
			BuildingNumber:   "6711",
			Community:        "University Village",
			Postcode:         "48722-7385",
			ZipCode:          "70123-4233",
			TimeZone:         "Asia/Novosibirsk",
			StreetSuffix:     "Flats",
			CitySuffix:       "land",
			CityPrefix:       "East",
			State:            "Arkansas",
			StateAbbr:        "NC",
			Country:          "Greenland",
			CountryCode:      "MG",
			CountryLongCode:  "ALA",
			FullAddress:      "Apt. 476 3356 Exie Dam, Crooksside, PA 20365",
		},
	}

	for _, test := range tests {
		random.Seed(1000)
		bundle.SetLocale(test.locale)
		suite.Equal(test.City, City(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.StreetName, StreetName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.StreetAddress, StreetAddress(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.SecondaryAddress, SecondaryAddress(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.BuildingNumber, BuildingNumber(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.Community, Community(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.Postcode, Postcode(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.ZipCode, ZipCode(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.TimeZone, TimeZone(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.StreetSuffix, StreetSuffix(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.CitySuffix, CitySuffix(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.CityPrefix, CityPrefix(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.State, State(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.StateAbbr, StateAbbr(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.Country, Country(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.CountryCode, CountryCode(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.CountryLongCode, CountryLongCode(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.FullAddress, FullAddress(), fmt.Sprintf("case: %s", test.locale))
	}
}
