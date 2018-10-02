package phonenumber

import (
	"github.com/ZhaoLion/faker/random"
)

// FakePhone a random phone
type FakePhone struct {
	PhoneNumber struct {
		Formats []string
	} `yaml:"phone_number"`
	CellPhone struct {
		Formats []string
	} `yaml:"cell_phone"`
}

// RandomPhoneNumber return random format phone number
func (p *FakePhone) RandomPhoneNumber() string {
	return random.Numerify(random.PickString(p.PhoneNumber.Formats))
}

// RandomCellPhone return random format cell number
func (p *FakePhone) RandomCellPhone() string {
	return random.Numerify(random.PickString(p.CellPhone.Formats))
}
