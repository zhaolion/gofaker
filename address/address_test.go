package address

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaolion/gofaker/random"
	"github.com/zhaolion/gofaker/test/mock"
	"gopkg.in/yaml.v2"
)

var bytes = []byte(`
city_prefix: [North, East, West, South, New, Lake, Port]
city_suffix: [town, ton, land, ville, berg, burgh, borough]
country: [Afghanistan, Albania, Algeria, American Samoa]
country_code: ["AD", "AE", "AF", "AG", "AI"]
country_code_long: ["ABW", "AFG", "AGO", "AIA", "ALA"]
building_number: ['#####', '####', '###']
community_prefix: [Park, Summer, Autumn, Paradise, Eagle, Pine, Royal, University, Willow]
community_suffix: [Village, Creek, Place, Pointe, Square, Oaks]
community:
  - "{{CommunityPrefix}} {{CommunitySuffix}}"
street_suffix: [Alley, Avenue, Branch, Bridge, Brook]
secondary_address: ['Apt. ###', 'Suite ###']
postcode: ['#####', '#####-####']
postcode_by_state:
  AL: '350##'
  AK: '995##'
  AS: '967##'
  AZ: '850##'
  AR: '717##'
state: [Alabama, Alaska, Arizona, Arkansas]
state_abbr: [AL, AK, AZ, AR, CA, CO, CT, DE]
time_zone: [Pacific/Midway, Pacific/Pago_Pago, Pacific/Honolulu, America/Juneau]
city:
  - "{{CityPrefix}} {{RandomFirstName}}{{CitySuffix}}"
  - "{{CityPrefix}} {{RandomFirstName}}"
  - "{{RandomFirstName}}{{CitySuffix}}"
  - "{{RandomLastName}}{{CitySuffix}}"
street_name:
  - "{{RandomFirstName}} {{StreetSuffix}}"
  - "{{RandomLastName}} {{StreetSuffix}}"
street_address:
  - "{{BuildingNumber}} {{StreetName}}"
full_address:
  - "{{StreetAddress}}, {{City}}, {{StateAbbr}} {{ZipCode}}"
  - "{{SecondaryAddress}} {{StreetAddress}}, {{City}}, {{StateAbbr}} {{ZipCode}}"
default_country: [United States of America]  
`)

var fakeAddress *FakeAddress

func init() {
	fakeAddress = NewFakeAddress(mock.NewFakeNameMock())
	err := yaml.Unmarshal(bytes, fakeAddress)
	if err != nil {
		log.Fatalf("unmarshal err: %s", err)
	}
}

func TestFakeAddressFuncs(t *testing.T) {
	random.Seed(1000)

	assert.Equal(t, "Abbie", fakeAddress.RandomFirstName())
	assert.Equal(t, "Abernathy", fakeAddress.RandomLastName())
	assert.Equal(t, "East", fakeAddress.CityPrefix())
	assert.Equal(t, "berg", fakeAddress.CitySuffix())
	assert.Equal(t, "Afghanistan", fakeAddress.Country())
	assert.Equal(t, "AE", fakeAddress.CountryCode())
	assert.Equal(t, "AIA", fakeAddress.CountryLongCode())
	assert.Equal(t, "055", fakeAddress.BuildingNumber())
	assert.Equal(t, "Summer", fakeAddress.CommunityPrefix())
	assert.Equal(t, "Village", fakeAddress.CommunitySuffix())
	assert.Equal(t, "Willow Place", fakeAddress.Community())
	assert.Equal(t, "Avenue", fakeAddress.StreetSuffix())
	assert.Equal(t, "Suite 670", fakeAddress.SecondaryAddress())
	assert.Equal(t, "04671", fakeAddress.Postcode())
	assert.Equal(t, "35017", fakeAddress.StatePostcode("AL"))
	assert.Equal(t, "Alabama", fakeAddress.State())
	assert.Equal(t, "DE", fakeAddress.StateAbbr())
	assert.Equal(t, "America/Juneau", fakeAddress.TimeZone())
	assert.Equal(t, "Abbieton", fakeAddress.City())
	assert.Equal(t, "Abernathy Alley", fakeAddress.StreetName())
	assert.Equal(t, "517 Abernathy Alley", fakeAddress.StreetAddress())
	assert.Equal(t, "2333 Abbey Avenue, Abernathyborough, CA 63476", fakeAddress.FullAddress())
	assert.Equal(t, "United States of America", fakeAddress.DefaultCountry())
}
