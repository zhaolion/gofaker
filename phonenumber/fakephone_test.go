package phonenumber

import (
	"testing"

	"github.com/ZhaoLion/faker/random"

	"gopkg.in/yaml.v2"
)

var fakePhoneBytes = []byte(`
phone_number:
  formats: 
    - '###-###-####'
    - '(###) ###-####'
    - '1-###-###-####'
    - '###.###.#### x#####'
cell_phone:
  formats:
    - '###-###-####'
    - '(###) ###-####'
    - '1-###-###-####'
    - '###.###.####'
`)

func TestFakePhone(t *testing.T) {
	random.Seed(110)

	phone := &FakePhone{}
	err := yaml.Unmarshal(fakePhoneBytes, phone)
	if err != nil {
		t.Fatalf("new FakePhone err: %s", err)
	}

	if phone.RandomPhoneNumber() != "1-163-061-5813" {
		t.Error("RandomPhoneNumber() should return \"1-163-061-5813\"")
	}
	if phone.RandomCellPhone() != "1-872-123-7517" {
		t.Error("RandomPhoneNumber() should return \"1-872-123-7517\"")
	}
}
