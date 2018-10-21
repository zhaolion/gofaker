package mock

import (
	"log"

	"github.com/zhaolion/gofaker/name"
	"gopkg.in/yaml.v2"
)

// NewFakeNameMock mock fake name
func NewFakeNameMock() *name.FakeName {
	fakeName := &name.FakeName{}
	err := yaml.Unmarshal(fakeNamebytes, fakeName)
	if err != nil {
		log.Fatalf("unmarshal FakeName err: %s", err)
	}

	return fakeName
}

var fakeNamebytes = []byte(`
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
man_last_name: [Antal, Babka, Bahna]
woman_last_name: [Antalová, Babková, Bahnová]
name:
  - "{{Prefix}} {{FirstName}} {{LastName}} {{Suffix}}"
name_with_middle:
  - "{{Prefix}} {{FirstName}} {{LastName}} {{LastName}}"
`)
