package faker

import (
	"github.com/zhaolion/faker/name"
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
