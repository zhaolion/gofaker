package faker

import "github.com/ZhaoLion/faker/phonenumber"

// RandomCellPhone random format cell phone
func RandomCellPhone() string {
	return currentBackend().Phone.RandomCellPhone()
}

// RandomPhoneNumber random format phone number
func RandomPhoneNumber() string {
	return currentBackend().Phone.RandomPhoneNumber()
}

// Phone fake phone
func Phone() *phonenumber.FakePhone {
	return currentBackend().Phone
}
