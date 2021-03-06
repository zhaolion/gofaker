package phonenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaolion/gofaker/random"
	"gopkg.in/yaml.v2"
)

var fakePhoneBytes = []byte(`
phone_number:
  formats: 
    - '###-###-####'
    - '(###) ###-####'
    - '1-###-###-####'
    - '###.###.#### x#####'
cell_phone:
  formats:
    - '###-###-####'
    - '(###) ###-####'
    - '1-###-###-####'
    - '###.###.####'
area_code: ["201", "202", "203"]
exchange_code: ["201", "202", "203"]
country_code: ["+1", "+2", "+3"]
lada_dos: ["33", "55"]
lada_tres: ["222", "223"]
`)

func TestFakePhone(t *testing.T) {
	random.Seed(110)

	phone := &FakePhone{}
	err := yaml.Unmarshal(fakePhoneBytes, phone)
	if err != nil {
		t.Fatalf("new FakePhone err: %s", err)
	}

	assert.Equal(t, 3, len(phone.AreaCodes))
	assert.Equal(t, 3, len(phone.ExchangeCodes))
	assert.Equal(t, 3, len(phone.CountryCodes))
	assert.Equal(t, "1-163-061-5813", phone.RandomPhoneNumber())
	assert.Equal(t, "1-872-123-7517", phone.RandomCellPhone())
	assert.Equal(t, "203", phone.AreaCode())
	assert.Equal(t, "202", phone.ExchangeCode())
	assert.Equal(t, "+3", phone.CountryCode())

	phone.AreaCodes = []string{"201"}
	phone.ExchangeCodes = []string{"202"}
	phone.CountryCodes = []string{"+1"}
	phone.PhoneNumber.Formats = []string{"{{AreaCode}}"}
	assert.Equal(t, "201", phone.RandomPhoneNumber())
	phone.PhoneNumber.Formats = []string{"({{CountryCode}})-{{AreaCode}}"}
	assert.Equal(t, "(+1)-201", phone.RandomPhoneNumber())
	phone.PhoneNumber.Formats = []string{"({{CountryCode}})-{{AreaCode}}-{{ExchangeCode}}"}
	assert.Equal(t, "(+1)-201-202", phone.RandomPhoneNumber())
	phone.PhoneNumber.Formats = []string{"({{CountryCode}})-{{AreaCode}}-{{ExchangeCode}}-{{Extension}}"}
	assert.Equal(t, "(+1)-201-202-0031", phone.RandomPhoneNumber())
	phone.PhoneNumber.Formats = []string{"({{CountryCode}})-{{SubscriberNumber}}-{{Extension}}"}
	assert.Equal(t, "(+1)-4496-5796", phone.RandomPhoneNumber())

	assert.Equal(t, 2, len(phone.LadaTress))
	assert.Equal(t, 2, len(phone.LadaDoses))
	assert.Equal(t, "222", phone.LadaTres())
	assert.Equal(t, "33", phone.LadaDos())
	phone.PhoneNumber.Formats = []string{"({{LadaTres}})-{{LadaDos}}-{{Extension}}"}
	assert.Equal(t, "(223)-55-0580", phone.RandomPhoneNumber())
	assert.Equal(t, "43", phone.FixedSubscriberNumber(2))
}
