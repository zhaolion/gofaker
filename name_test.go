package faker

import (
	"fmt"

	"github.com/zhaolion/faker/random"
)

func (suite *Suite) TestName() {
	var tests = []struct {
		locale          string
		Name            string
		NameWithMiddle  string
		FirstName       string
		MiddleName      string
		LastName        string
		Prefix          string
		Suffix          string
		MaleFirstName   string
		FemaleFirstName string
	}{
		{
			locale:          "en",
			Name:            "Truman Wintheiser III",
			NameWithMiddle:  "Elijah Kub Swaniawski",
			MiddleName:      "Bartoletti",
			FirstName:       "Dallas",
			LastName:        "Paucek",
			Prefix:          "Mrs.",
			Suffix:          "Sr.",
			MaleFirstName:   "Wilford",
			FemaleFirstName: "Carrie",
		},
		{
			locale:          "bg",
			Name:            "Кирова Мария",
			NameWithMiddle:  "Иванка Reichel Osinski",
			MiddleName:      "Ryan",
			FirstName:       "Стоянка",
			LastName:        "O'Conner",
			Prefix:          "Mrs.",
			Suffix:          "III",
			MaleFirstName:   "Петър",
			FemaleFirstName: "Румяна",
		},
		{
			locale:          "ca",
			Name:            "Sr. Gerard Adell Roma",
			NameWithMiddle:  "Aleu Amengual Guarch",
			MiddleName:      "Roig",
			FirstName:       "Eudald",
			LastName:        "Danès",
			Prefix:          "Mrs.",
			Suffix:          "DDS",
			MaleFirstName:   "Josefí",
			FemaleFirstName: "Benvinguda",
		},
		{
			locale:          "ca-CAT",
			Name:            "Truman Wintheiser III",
			NameWithMiddle:  "Elijah Kub Swaniawski",
			MiddleName:      "Bartoletti",
			FirstName:       "Dallas",
			LastName:        "Paucek",
			Prefix:          "Mrs.",
			Suffix:          "Sr.",
			MaleFirstName:   "Wilford",
			FemaleFirstName: "Carrie",
		},
		{
			locale:          "da-DK",
			Name:            "Cand.jur. Peter Rasmussen",
			NameWithMiddle:  "Kirsten Pedersen Larsen",
			FirstName:       "Henrik",
			LastName:        "Christensen",
			Prefix:          "Cand.jur.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
			MiddleName:      "Hansen",
		},
		{
			locale:          "de-AT",
			Name:            "Ivan zu Greschner",
			NameWithMiddle:  "Lino Brandis Holinski",
			FirstName:       "Anne",
			LastName:        "Edorh",
			Prefix:          "Prof. Dr.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
			MiddleName:      "Badane",
		},
		{
			locale:          "de-CH",
			Name:            "Truman Wintheiser III",
			NameWithMiddle:  "Elijah Kub Swaniawski",
			MiddleName:      "Bartoletti",
			FirstName:       "Dallas",
			LastName:        "Paucek",
			Prefix:          "Mrs.",
			Suffix:          "Sr.",
			MaleFirstName:   "Wilford",
			FemaleFirstName: "Carrie",
		},
		{
			locale:          "de",
			Name:            "Toni Dolch",
			NameWithMiddle:  "Lisa Gerschler Birkemeyer",
			MiddleName:      "Hirt",
			FirstName:       "Henri",
			LastName:        "Geisler",
			Prefix:          "Fr.",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "ee",
			Name:            "Mare Vaino",
			NameWithMiddle:  "Evert Raag Jänes",
			MiddleName:      "Raud",
			FirstName:       "Aili",
			LastName:        "Aasmäe",
			Prefix:          "Mr.",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "en-au-ocker",
			Name:            "Cooper Smith IV",
			NameWithMiddle:  "Isabella Walker Harris",
			MiddleName:      "Harris",
			FirstName:       "Chloe",
			LastName:        "Taylor",
			Prefix:          "Mrs.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "en-AU",
			Name:            "Mackenzie Dooley IV",
			NameWithMiddle:  "Zachary Wolf Goodwin",
			MiddleName:      "Muller",
			FirstName:       "Elijah",
			LastName:        "Erdman",
			Prefix:          "Mrs.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "en-BORK",
			Name:            "Truman Wintheiser III",
			NameWithMiddle:  "Elijah Kub Swaniawski",
			MiddleName:      "Bartoletti",
			FirstName:       "Dallas",
			LastName:        "Paucek",
			Prefix:          "Mrs.",
			Suffix:          "Sr.",
			MaleFirstName:   "Wilford",
			FemaleFirstName: "Carrie",
		},
	}

	for _, test := range tests {
		random.Seed(1000)
		bundle.SetLocale(test.locale)
		suite.Equal(test.Name, Name(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.NameWithMiddle, NameWithMiddle(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.FirstName, FirstName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.LastName, LastName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.Prefix, PrefixName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.Suffix, SuffixName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.MaleFirstName, MaleFirstName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.FemaleFirstName, FemaleFirstName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.MiddleName, MiddleName(), fmt.Sprintf("case: %s", test.locale))
	}
}
