package faker

func (suite *Suite) TestNewLocaleBackend() {
	enBackend, err := newBackend(localeFilePath("en"))
	suite.Nil(err)
	suite.NotNil(enBackend)

	target, err := NewLocaleBackend("de-AT")
	suite.Nil(err)
	suite.NotNil(target)

	suite.NotEqual(enBackend.Phone, target.Phone)
	suite.Equal(enBackend.Name, target.Name)

	ignoreLocales := []string{
		"en",
		"ca",
		"en-BORK",
		"fa",
		"no",
	}

	for _, locale := range Locales {
		if containString(ignoreLocales, locale) {
			continue
		}

		target, err := NewLocaleBackend(locale)
		suite.Nil(err)
		suite.NotNil(target)
		suite.Equal(enBackend.Name, target.Name, locale)
	}
}
