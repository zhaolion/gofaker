package random

func (suite *Suite) TestFixedAlpha() {
	suite.Equal("HdaI", FixedAlpha(4))
}

func (suite *Suite) TestFixedNumeric() {
	suite.Equal("5108", FixedNumeric(4))
}

func (suite *Suite) TestFixedAlphaNumeric() {
	suite.Equal("VJe0", FixedAlphaNumeric(4))
}

func (suite *Suite) TestNumerify() {
	suite.Equal("416", Numerify("###"))
}

func (suite *Suite) TestLetterify() {
	suite.Equal("HDA", Letterify("???"))
}

func (suite *Suite) TestBothify() {
	suite.Equal("416IQV", Bothify("###???"))
}
