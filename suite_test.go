package gofaker

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zhaolion/gofaker/random"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type Suite struct {
	suite.Suite
}

// The SetupTest method will be run before every test in the suite.
func (suite *Suite) SetupTest() {
	random.Seed(1000)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
