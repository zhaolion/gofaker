package gofaker

func (suite *Suite) TestSetLocale() {
	SetLocale("zh-CN")
	suite.Equal("zh-CN", bundle.currentLocale)
	backend, ok := bundle.backends["zh-CN"]
	suite.True(ok)
	suite.Equal(backend, bundle.currentBackend)
}
