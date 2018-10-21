package gofaker

import (
	"fmt"
)

func (suite *Suite) TestNewBundle() {
	bundle := NewBundle()
	suite.Equal(len(bundle.locales), len(bundle.backends))
	for name, backend := range bundle.backends {
		suite.NotNil(backend, fmt.Sprintf("load bundle backends[%s] error", name))
	}
}
