package random

import (
	"testing"
)

func TestNumber(t *testing.T) {
	if Number(10, 20) != 17 {
		t.Error("Number(10, 20) should return 17")
	}
}

func TestNumerify(t *testing.T) {
	if Numerify("###") != "816" {
		t.Error("Numerify(###) should return 816")
	}
}
