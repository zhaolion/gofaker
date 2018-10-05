package random

func (suite *Suite) TestPickString() {
	suite.Equal("", PickString([]string{}))
	suite.Equal("2", PickString([]string{"1", "2", "3"}))
}
