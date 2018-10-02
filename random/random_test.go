package random

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Seed(110)
	code := m.Run()
	os.Exit(code)
}
