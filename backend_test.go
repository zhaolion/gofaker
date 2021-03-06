package gofaker

func (suite *Suite) TestNewLocaleBackend() {
	enBackend, err := initDefaultBackend()
	suite.Nil(err)
	suite.NotNil(enBackend)

	target, err := NewLocaleBackend("ca-CAT")
	suite.Nil(err)
	suite.NotNil(target)

	suite.NotEqual(enBackend.Phone, target.Phone)
	suite.Equal(enBackend.Name, target.Name)
}
