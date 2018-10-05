package random

func (suite *Suite) TestNumber() {
	suite.Equal(19, Number(10, 20))
}
