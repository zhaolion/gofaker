package name

import (
	"testing"

	"log"

	"github.com/stretchr/testify/assert"
	"github.com/zhaolion/faker/random"
	"gopkg.in/yaml.v2"
)

var bytes = []byte(`
male_first_name: [Aaron, Abdul, Abe, Abel]
female_first_name: [Abbey, Abbie, Abby]
first_name:
  - "{{FemaleFirstName}}"
last_name: [Abbott, Abernathy, Abshire]
prefix: [Mr., Mrs., Ms., Miss, Dr.]
suffix: [Jr., Sr., I, II, III, IV, V, MD, DDS, PhD, DVM]
name:
  - "{{Prefix}} {{FirstName}} {{LastName}} {{Suffix}}"
name_with_middle:
  - "{{Prefix}} {{FirstName}} {{LastName}} {{LastName}}"
`)

var fakeName *FakeName

func init() {
	fakeName = &FakeName{}
	err := yaml.Unmarshal(bytes, fakeName)
	if err != nil {
		log.Fatalf("unmarshal err: %s", err)
	}
}

func TestFakeNameFuncs(t *testing.T) {
	random.Seed(1000)

	assert.Equal(t, "Abdul", fakeName.MaleFirstName())
	assert.Equal(t, "Abbie", fakeName.FemaleFirstName())
	assert.Equal(t, "Mr.", fakeName.Prefix())
	assert.Equal(t, "IV", fakeName.Suffix())
	assert.Equal(t, "Abbie", fakeName.FirstName())
	assert.Equal(t, "Mrs. Abbie Abshire DDS", fakeName.Name())
	assert.Equal(t, "Mr. Abby Abernathy Abbott", fakeName.NameWithMiddle())
}
