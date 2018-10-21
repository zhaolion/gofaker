package gofaker

import (
	"github.com/zhaolion/gofaker/name"
)

// FakeName return fake name
func FakeName() *name.FakeName {
	if currentBackend().Name == nil {
		return defaultBackend().Name
	}

	return currentBackend().Name
}

// Name name of person
func Name() string {
	if FakeName() == nil {
		return ""
	}

	return FakeName().Name()
}

// NameWithMiddle name with middle
func NameWithMiddle() string {
	if FakeName() == nil {
		return ""
	}

	return FakeName().NameWithMiddle()
}

// FirstName first name
func FirstName() string {
	if FakeName() == nil {
		return ""
	}

	return FakeName().FirstName()
}

// MiddleName middle name
func MiddleName() string {
	if FakeName() == nil {
		return ""
	}

	return FakeName().MiddleName()
}

// LastName last name
func LastName() string {
	if FakeName() == nil {
		return ""
	}

	return FakeName().LastName()
}

// PrefixName prefix of name
func PrefixName() string {
	if FakeName() == nil {
		return ""
	}

	return FakeName().Prefix()
}

// SuffixName suffix of name
func SuffixName() string {
	if FakeName() == nil {
		return ""
	}

	return FakeName().Suffix()
}

// MaleFirstName male first name
func MaleFirstName() string {
	if FakeName() == nil {
		return ""
	}

	return FakeName().MaleFirstName()
}

// FemaleFirstName female first name
func FemaleFirstName() string {
	if FakeName() == nil {
		return ""
	}

	return FakeName().FemaleFirstName()
}
