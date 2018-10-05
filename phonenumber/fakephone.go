package phonenumber

import (
	"github.com/ZhaoLion/faker/random"
	"github.com/hoisie/mustache"
)

// FakePhone a random phone
type FakePhone struct {
	PhoneNumber struct {
		Formats []string
	} `yaml:"phone_number,flow"`
	CellPhone struct {
		Formats []string
	} `yaml:"cell_phone,flow"`
	AreaCodes     []string `yaml:"area_code,flow"`
	ExchangeCodes []string `yaml:"exchange_code,flow"`
	CountryCodes   []string `yaml:"country_code,flow"`
}

// RandomPhoneNumber return random format phone number
func (p *FakePhone) RandomPhoneNumber() string {
	number := random.PickString(p.PhoneNumber.Formats)
	return random.Numerify(mustache.Render(number, p))
}

// RandomCellPhone return random format cell number
func (p *FakePhone) RandomCellPhone() string {
	number := random.PickString(p.CellPhone.Formats)
	return random.Numerify(mustache.Render(number, p))
}

// AreaCode random AreaCode
func (p *FakePhone) AreaCode() string {
	return random.PickString(p.AreaCodes)
}

// ExchangeCode random ExchangeCode
func (p *FakePhone) ExchangeCode() string {
	return random.PickString(p.ExchangeCodes)
}

// CountryCode random CountryCode
func (p *FakePhone) CountryCode() string {
	return random.PickString(p.CountryCodes)
}
