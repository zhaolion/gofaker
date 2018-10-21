package faker

import (
	"fmt"

	"github.com/zhaolion/gofaker/random"
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
		{
			locale:          "en-CA",
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
			locale:          "en-GB",
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
			locale:          "en-IND",
			Name:            "Brajesh Jha IV",
			NameWithMiddle:  "Chaturbhuj Joshi Dwivedi",
			MiddleName:      "Chaturvedi",
			FirstName:       "Sharda",
			LastName:        "Naik",
			Prefix:          "Mrs.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "en-MS",
			Name:            "Datin Siti binti Abu Bakar",
			NameWithMiddle:  "Yun Xuan Osinski Kub",
			MiddleName:      "Prosacco",
			FirstName:       "Farah",
			LastName:        "Howe",
			Prefix:          "Dato",
			Suffix:          "DDS",
			MaleFirstName:   "Hugh",
			FemaleFirstName: "Dorine",
		},
		{
			locale:          "en-NEP",
			Name:            "Kamal Adhikari Chettri",
			NameWithMiddle:  "Sunita Maharjan Shrestha",
			FirstName:       "Khem",
			MiddleName:      "Dongol",
			LastName:        "Khadka",
			Prefix:          "Mrs.",
			Suffix:          "V",
			MaleFirstName:   "Himal",
			FemaleFirstName: "Griahma",
		},
		{
			locale:          "en-NG",
			Name:            "Omawunmi Adegoke",
			NameWithMiddle:  "Miss Chinedu Onyinyechukwu Maryjane",
			FirstName:       "Abisoye",
			MiddleName:      "Olumide",
			LastName:        "Mustapha",
			Prefix:          "Mr.",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "en-NZ",
			Name:            "Grace Dooley IV",
			NameWithMiddle:  "Tane Wolf Goodwin",
			FirstName:       "Aroha",
			MiddleName:      "Muller",
			LastName:        "Erdman",
			Prefix:          "Mrs.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "en-PAK",
			Name:            "Waheed Zarnosh Khan IV",
			NameWithMiddle:  "Munawar Ali Raza Iqbal",
			FirstName:       "Muzaffar",
			MiddleName:      "A Haq Ansari",
			LastName:        "Sindho",
			Prefix:          "Mrs.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "en-SG",
			Name:            "Terence Tan Kah Seng",
			NameWithMiddle:  "You Ming Tay Au",
			FirstName:       "Wei Jie",
			MiddleName:      "Chow",
			LastName:        "Wong",
			Prefix:          "Mrs.",
			Suffix:          "V",
			MaleFirstName:   "Jun Jie",
			FemaleFirstName: "Jing Ning",
		},
		{
			locale:          "en-UG",
			Name:            "Truman Bakabulindi III",
			NameWithMiddle:  "Elijah Kasekende Bwana",
			MiddleName:      "Okumu",
			FirstName:       "Dallas",
			LastName:        "Byanyima",
			Prefix:          "Mrs.",
			Suffix:          "Sr.",
			MaleFirstName:   "Wilford",
			FemaleFirstName: "Carrie",
		},
		{
			locale:          "en-US",
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
			locale:          "en-ZA",
			Name:            "Cameron Kay IV",
			NameWithMiddle:  "Ofentse Mabunda Keyser",
			MiddleName:      "Niemann",
			FirstName:       "Kimberley",
			LastName:        "Jackson",
			Prefix:          "Mrs.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "es-MX",
			Name:            "Marina Hernández Soriano",
			NameWithMiddle:  "Ignacio Mercado Yáñez",
			MiddleName:      "Sandoval",
			FirstName:       "Pablo",
			LastName:        "Carrillo",
			Prefix:          "Sra.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "es",
			Name:            "Sra.DoloresRíosTafoya",
			NameWithMiddle:  "Elisa Juárez Reyna V",
			MiddleName:      "Pelayo",
			FirstName:       "Caridad",
			LastName:        "Corona",
			Prefix:          "Sra.",
			Suffix:          "I",
			MaleFirstName:   "Felipe",
			FemaleFirstName: "Barbara",
		},
		{
			locale:          "fa",
			Name:            "سورن ملک IV",
			NameWithMiddle:  "رادبانو علی\u200cآبادی نامور",
			MiddleName:      "مصباح\u200cزاده",
			FirstName:       "فرهود",
			LastName:        "عبدالملکی",
			Prefix:          "دکتر",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "fi-FI",
			Name:            "Pirkko Koivisto",
			NameWithMiddle:  "Miss Timo Virtanen Laaksonen",
			MiddleName:      "Virtanen",
			FirstName:       "Hannele",
			LastName:        "Hirvonen",
			Prefix:          "Mr.",
			Suffix:          "DVM",
			MaleFirstName:   "Aleksi",
			FemaleFirstName: "Anne",
		},
		{
			locale:          "fr-CA",
			Name:            "Mme Pierre Rénault",
			NameWithMiddle:  "Théo Clemént Lemairé",
			MiddleName:      "Roy",
			FirstName:       "Juliette",
			LastName:        "Lucas",
			Prefix:          "Mme",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "fr-CH",
			Name:            "Arthur Dumont",
			NameWithMiddle:  "Enzo Marchand Charlés",
			MiddleName:      "Blanchard",
			FirstName:       "Quentin",
			LastName:        "Rivieré",
			Prefix:          "M",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "fr",
			Name:            "Arthur Martin",
			NameWithMiddle:  "Enzo Rodriguéz Clemént",
			MiddleName:      "Rodriguéz",
			FirstName:       "Quentin",
			LastName:        "Schneidér",
			Prefix:          "M",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "he",
			Name:            "אליעזר פרץ",
			NameWithMiddle:  "Miss אמונה שטרן משה",
			MiddleName:      "שלום",
			FirstName:       "אפרים",
			LastName:        "סויסה",
			Prefix:          "Mr.",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "id",
			Name:            "Vira Riyadi",
			NameWithMiddle:  "Miss Anton Darmadi Yulianto",
			MiddleName:      "Zaenal",
			FirstName:       "Tomi",
			LastName:        "Unggul",
			Prefix:          "Mr.",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "it",
			Name:            "Eustachio Amato",
			NameWithMiddle:  "Valdo Lombardo Silvestri",
			MiddleName:      "De rosa",
			FirstName:       "Arcibaldo",
			LastName:        "Giuliani",
			Prefix:          "Dott.",
			Suffix:          "",
			MaleFirstName:   "Odell",
			FemaleFirstName: "Siu",
		},
		{
			locale:          "ja",
			Name:            "藤原 優花",
			NameWithMiddle:  "Miss 陽子 中川 岩崎",
			MiddleName:      "中川",
			FirstName:       "大翔",
			LastName:        "西村",
			Prefix:          "Mr.",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "ko",
			Name:            "강 하은",
			NameWithMiddle:  "Miss 우진 한 서",
			MiddleName:      "한",
			FirstName:       "서현",
			LastName:        "정",
			Prefix:          "Mr.",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "lv",
			Name:            "Pāvels Balodis",
			NameWithMiddle:  "Edīte Jansons Vasiļjevs",
			MiddleName:      "Muižnieks",
			FirstName:       "Ainis",
			LastName:        "Āboliņš",
			Prefix:          "Mr.",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "nb-NO",
			Name:            "Sindre Johansen III",
			NameWithMiddle:  "Mathias Nordskaug Solberg",
			MiddleName:      "Aasen",
			FirstName:       "Sunniva",
			LastName:        "Berntsen",
			Prefix:          "Prof.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "nl",
			Name:            "Niels Bakker III",
			NameWithMiddle:  "Daniël Visser Brink",
			MiddleName:      "Dijk",
			FirstName:       "Nick",
			LastName:        "Groot",
			Prefix:          "Mevr. Dr.",
			Suffix:          "V",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "pl",
			Name:            "Angelina Młynarczyk",
			NameWithMiddle:  "Julita Kujawa Dróżdż",
			MiddleName:      "Piechota",
			FirstName:       "Weronika",
			LastName:        "Kujawski",
			Prefix:          "Pani",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "pt-BR",
			Name:            "Isaac Martins Neto",
			NameWithMiddle:  "João Gabriel Fogaça Corte",
			MiddleName:      "Ambrósio",
			FirstName:       "Vinícius",
			LastName:        "Castelo",
			Prefix:          "Sra.",
			Suffix:          "Filho",
			MaleFirstName:   "Vitor Hugo",
			FemaleFirstName: "Dalila",
		},
		{
			locale:          "pt",
			Name:            "Sara Araújo Filho",
			NameWithMiddle:  "Silas Melo Santos",
			MiddleName:      "Batista",
			FirstName:       "Rafael",
			LastName:        "Costa",
			Prefix:          "Dr.",
			Suffix:          "Neto",
			MaleFirstName:   "Rupert",
			FemaleFirstName: "Bettina",
		},
		{
			locale:          "ru",
			Name:            "Сорокина Анжела",
			NameWithMiddle:  "Анфиса Селиверстов Пестова",
			MiddleName:      "Лобанова",
			FirstName:       "Арсений",
			LastName:        "Орлова",
			Prefix:          "Mrs.",
			Suffix:          "Sr.",
			MaleFirstName:   "Лука",
			FemaleFirstName: "Алла",
		},
		{
			locale:          "sk",
			Name:            "Metod Líška",
			NameWithMiddle:  "Július Kaliská Halák",
			MiddleName:      "Tomko",
			FirstName:       "Liana",
			LastName:        "Cyprich",
			Prefix:          "JUDr.",
			Suffix:          "Phd.",
			MaleFirstName:   "Dalibor",
			FemaleFirstName: "Petronela",
		},
		{
			locale:          "sv",
			Name:            "civ.ing. Žaneta Nilsson",
			NameWithMiddle:  "Aurel Andersson Persson",
			MiddleName:      "Andersson",
			FirstName:       "Timotej",
			LastName:        "Olsson",
			Prefix:          "fil.mag.",
			Suffix:          "Sr.",
			MaleFirstName:   "Dalibor",
			FemaleFirstName: "Petronela",
		},
		{
			locale:          "tr",
			Name:            "Caner Sağdıç",
			NameWithMiddle:  "İrem Ekkaldır Zengel",
			MiddleName:      "Zengel",
			FirstName:       "Selim",
			LastName:        "Sağdıç",
			Prefix:          "Dr.",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
		},
		{
			locale:          "uk",
			Name:            "Ющук Броніслава",
			NameWithMiddle:  "Анастасій Каськів Лученко",
			MiddleName:      "Федоришин",
			FirstName:       "Лаврін",
			LastName:        "Михайлюк",
			Prefix:          "Dr.",
			Suffix:          "IV",
			MaleFirstName:   "Тимофій",
			FemaleFirstName: "Добромила",
		},
		{
			locale:          "vi",
			Name:            "Lê Dung Kim",
			NameWithMiddle:  "Đoàn Khoa Hữu",
			MiddleName:      "Nhàn",
			FirstName:       "Mai",
			LastName:        "Tiến",
			Prefix:          "Mr.",
			Suffix:          "DVM",
			MaleFirstName:   "Dallas",
			FemaleFirstName: "Assunta",
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
