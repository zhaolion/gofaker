package faker

// RandomCellPhone random format cell phone
func RandomCellPhone() string {
	return currentBackend().Phone.RandomCellPhone()
}

// RandomPhoneNumber random format phone number
func RandomPhoneNumber() string {
	return currentBackend().Phone.RandomPhoneNumber()
}
