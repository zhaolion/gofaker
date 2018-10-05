package faker

func (suite *Suite) TestNewBackend() {
	backend, err := NewBackend(dataFilePath("en.yml"))
	suite.Nil(err)
	suite.NotNil(backend)
}
