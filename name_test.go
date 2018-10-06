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
	}

	for _, test := range tests {
		random.Seed(1000)
		bundle.SetLocale(test.locale)
		suite.Equal(test.Name, Name(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.NameWithMiddle, NameWithMiddle(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.FirstName, FirstName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.LastName, LastName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.Prefix, Prefix(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.Suffix, Suffix(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.MaleFirstName, MaleFirstName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.FemaleFirstName, FemaleFirstName(), fmt.Sprintf("case: %s", test.locale))
		suite.Equal(test.MiddleName, MiddleName(), fmt.Sprintf("case: %s", test.locale))
	}
}
