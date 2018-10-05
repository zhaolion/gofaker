package phonenumber

import (
	"github.com/ZhaoLion/faker/random"
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
	return random.Numerify(random.PickString(p.PhoneNumber.Formats))
}

// RandomCellPhone return random format cell number
func (p *FakePhone) RandomCellPhone() string {
	return random.Numerify(random.PickString(p.CellPhone.Formats))
}

func (p *FakePhone) AreaCode() string {
	return random.Numerify(random.PickString(p.AreaCodes))
}

func (p *FakePhone) ExchangeCode() string {
	return random.Numerify(random.PickString(p.ExchangeCodes))
}

func (p *FakePhone) CountryCode() string {
	return random.Numerify(random.PickString(p.CountryCodes))
}
