package name

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaolion/gofaker/random"
	"gopkg.in/yaml.v2"
)

var bytes = []byte(`
male_first_name: [Aaron, Abdul, Abe, Abel]
male_middle_name: [Aaron, Abdul, Abe, Abel]
male_last_name: [Aaron, Abdul, Abe, Abel]
female_first_name: [Abbey, Abbie, Abby]
female_middle_name: [Abbey, Abbie, Abby]
female_last_name: [Abbey, Abbie, Abby]
first_name:
  - "{{FemaleFirstName}}"
last_name: [Abbott, Abernathy, Abshire]
prefix: [Mr., Mrs., Ms., Miss, Dr.]
male_prefix: [Mr., Mrs., Ms., Miss, Dr.]
female_prefix: [Mr., Mrs., Ms., Miss, Dr.]
suffix: [Jr., Sr., I, II, III, IV, V, MD, DDS, PhD, DVM]
nobility_title_prefix: [zu, von, vom, von der]
nobility_title: [Baron, Baronin, Freiherr, Freiherrin, Graf, Gräfin]
ocker_first_name: [Bazza, Bluey, Davo, Johno, Shano, Shazza, Wazza, Charl, Darl]
malay_male_first_name: [Abu, Ahmad, Malik]
malay_female_first_name: [Siti, Aminah, Aiza]
chinese_male_first_name: [Jin Quan, Wen Jun, Jun Jie]
chinese_male_last_name: [Tan, Lim, Lee, Ng, Ong, Wong]
chinese_female_first_name: [Xiu Yi, Wai Teng, Sing Yee]
male_english_name: [Leon, Bryan, Jack, Stephen, Andy]
female_english_name: [Alicia, Caitlin, Denise, Emerald]
feminine_name: ["Emma", "Sara", "Thea", "Ida"]
masculine_name: ["Markus", "Mathias", "Kristian"]
tussenvoegsel: ["van", "van de", "van den"]
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
	assert.Equal(t, "Abshire", fakeName.MiddleName())
	assert.Equal(t, "Abe", fakeName.MaleMiddleName())
	assert.Equal(t, "Abel", fakeName.MaleLastName())
	assert.Equal(t, "Abbie", fakeName.FemaleMiddleName())
	assert.Equal(t, "Abbey", fakeName.FemaleLastName())
	assert.Equal(t, "Mr.", fakeName.MalePrefix())
	assert.Equal(t, "Miss", fakeName.FemalePrefix())
	assert.Equal(t, "zu", fakeName.NobilityTitlePrefix())
	assert.Equal(t, "Freiherrin", fakeName.NobilityTitle())
	assert.Equal(t, "Shano", fakeName.OckerFirstName())
	assert.Equal(t, "Abu", fakeName.MalayMaleFirstName())
	assert.Equal(t, "Aminah", fakeName.MalayFemaleFirstName())
	assert.Equal(t, "Wen Jun", fakeName.ChineseMaleFirstName())
	assert.Equal(t, "Lim", fakeName.ChineseMaleLastName())
	assert.Equal(t, "Wai Teng", fakeName.ChineseFemaleFirstName())
	assert.Equal(t, "Leon", fakeName.MaleEnglishName())
	assert.Equal(t, "Emerald", fakeName.FemaleEnglishName())
	assert.Equal(t, "Ida", fakeName.FeminineName())
	assert.Equal(t, "Kristian", fakeName.MasculineName())
	assert.Equal(t, "van de", fakeName.Tussenvoegsel())
}
