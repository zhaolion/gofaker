package address

import (
	"github.com/hoisie/mustache"
	"github.com/zhaolion/gofaker/name"
	"github.com/zhaolion/gofaker/random"
)

// FakeAddress fake random address
type FakeAddress struct {
	CityPrefixs        []string          `yaml:"city_prefix,flow"`
	CitySuffixs        []string          `yaml:"city_suffix,flow"`
	Countries          []string          `yaml:"country,flow"`
	CountryCodes       []string          `yaml:"country_code,flow"`
	CountryLongCodes   []string          `yaml:"country_code_long,flow"`
	BuildingNumbers    []string          `yaml:"building_number,flow"`
	CommunityPrefixs   []string          `yaml:"community_prefix,flow"`
	CommunitySuffixs   []string          `yaml:"community_suffix,flow"`
	Communities        []string          `yaml:"community,flow"`
	StreetSuffixs      []string          `yaml:"street_suffix,flow"`
	StreetNames        []string          `yaml:"street_name,flow"`
	StreetAddresses    []string          `yaml:"street_address,flow"`
	SecondaryAddresses []string          `yaml:"secondary_address,flow"`
	Postcodes          []string          `yaml:"postcode,flow"`
	PostcodesByState   map[string]string `yaml:"postcode_by_state"`
	States             []string          `yaml:"state,flow"`
	StateAbbrs         []string          `yaml:"state_abbr,flow"`
	TimeZones          []string          `yaml:"time_zone,flow"`
	Cities             []string          `yaml:"city,flow"`
	FullAddresses      []string          `yaml:"full_address,flow"`
	DefaultCountries   []string          `yaml:"default_country,flow"`

	name *name.FakeName
}

// NewFakeAddress fake address
func NewFakeAddress(name *name.FakeName) *FakeAddress {
	return &FakeAddress{
		name: name,
	}
}

// RandomFirstName random name.FirstName
// inner render func
func (fa *FakeAddress) RandomFirstName() string {
	return fa.name.FirstName()
}

// RandomLastName random name.LastName
// inner render func
func (fa *FakeAddress) RandomLastName() string {
	return fa.name.LastName()
}

// CityPrefix city_prefix
func (fa *FakeAddress) CityPrefix() string {
	return random.PickString(fa.CityPrefixs)
}

// CitySuffix city_suffix
func (fa *FakeAddress) CitySuffix() string {
	return random.PickString(fa.CitySuffixs)
}

// Country country
func (fa *FakeAddress) Country() string {
	return random.PickString(fa.Countries)
}

// CountryCode country_code
func (fa *FakeAddress) CountryCode() string {
	return random.PickString(fa.CountryCodes)
}

// CountryLongCode country_code_long
func (fa *FakeAddress) CountryLongCode() string {
	return random.PickString(fa.CountryLongCodes)
}

// BuildingNumber building_number
func (fa *FakeAddress) BuildingNumber() string {
	number := random.PickString(fa.BuildingNumbers)
	return random.Numerify(number)
}

// CommunityPrefix community_prefix
func (fa *FakeAddress) CommunityPrefix() string {
	return random.PickString(fa.CommunityPrefixs)
}

// CommunitySuffix community_suffix
func (fa *FakeAddress) CommunitySuffix() string {
	return random.PickString(fa.CommunitySuffixs)
}

// Community community
func (fa *FakeAddress) Community() string {
	c := random.PickString(fa.Communities)
	return random.Numerify(mustache.Render(c, fa))
}

// StreetSuffix street_suffix
func (fa *FakeAddress) StreetSuffix() string {
	return random.PickString(fa.StreetSuffixs)
}

// SecondaryAddress secondary_address
func (fa *FakeAddress) SecondaryAddress() string {
	sa := random.PickString(fa.SecondaryAddresses)
	return random.Numerify(sa)
}

// Postcode postcode
func (fa *FakeAddress) Postcode() string {
	c := random.PickString(fa.Postcodes)
	return random.Numerify(c)
}

// StatePostcode postcode_by_state
func (fa *FakeAddress) StatePostcode(state string) string {
	c := fa.PostcodesByState[state]
	return random.Numerify(c)
}

// State state
func (fa *FakeAddress) State() string {
	return random.PickString(fa.States)
}

// StateAbbr state_abbr
func (fa *FakeAddress) StateAbbr() string {
	return random.PickString(fa.StateAbbrs)
}

// TimeZone time_zone
func (fa *FakeAddress) TimeZone() string {
	return random.PickString(fa.TimeZones)
}

// City city
func (fa *FakeAddress) City() string {
	city := random.PickString(fa.Cities)
	return random.Numerify(mustache.Render(city, fa))
}

// StreetName street_name
func (fa *FakeAddress) StreetName() string {
	sn := random.PickString(fa.StreetNames)
	return random.Numerify(mustache.Render(sn, fa))
}

// StreetAddress street_address
func (fa *FakeAddress) StreetAddress() string {
	sa := random.PickString(fa.StreetAddresses)
	return random.Numerify(mustache.Render(sa, fa))
}

// FullAddress full_address
func (fa *FakeAddress) FullAddress() string {
	addr := random.PickString(fa.FullAddresses)
	return random.Numerify(mustache.Render(addr, fa))
}

// DefaultCountry default_country
func (fa *FakeAddress) DefaultCountry() string {
	return random.PickString(fa.DefaultCountries)
}

// ZipCode zip_code
func (fa *FakeAddress) ZipCode() string {
	return fa.Postcode()
}
