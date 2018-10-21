package gofaker

import (
	"github.com/zhaolion/gofaker/address"
)

// FakeAddress current fake address
func FakeAddress() *address.FakeAddress {
	if currentBackend().Address == nil {
		return defaultBackend().Address
	}

	return currentBackend().Address
}

// City city
func City() string {
	return FakeAddress().City()
}

// StreetName street_name
func StreetName() string {
	return FakeAddress().StreetName()
}

// StreetAddress street_address
func StreetAddress() string {
	return FakeAddress().StreetAddress()
}

// SecondaryAddress secondary_address
func SecondaryAddress() string {
	return FakeAddress().SecondaryAddress()
}

// BuildingNumber building_number
func BuildingNumber() string {
	return FakeAddress().BuildingNumber()
}

// Community community
func Community() string {
	return FakeAddress().Community()
}

// Postcode postcode
func Postcode() string {
	return FakeAddress().Postcode()
}

// ZipCode zip_code
func ZipCode() string {
	return FakeAddress().ZipCode()
}

// TimeZone time_zone
func TimeZone() string {
	return FakeAddress().TimeZone()
}

// StreetSuffix street_suffix
func StreetSuffix() string {
	return FakeAddress().StreetSuffix()
}

// CitySuffix city_suffix
func CitySuffix() string {
	return FakeAddress().CitySuffix()
}

// CityPrefix city_prefix
func CityPrefix() string {
	return FakeAddress().CityPrefix()
}

// State state
func State() string {
	return FakeAddress().State()
}

// StateAbbr state_abbr
func StateAbbr() string {
	return FakeAddress().StateAbbr()
}

// Country country
func Country() string {
	return FakeAddress().Country()
}

// CountryCode country_code
func CountryCode() string {
	return FakeAddress().CountryCode()
}

// CountryLongCode country_long_code
func CountryLongCode() string {
	return FakeAddress().CountryLongCode()
}

// FullAddress full_address
func FullAddress() string {
	return FakeAddress().FullAddress()
}
