package random

func (suite *Suite) TestNumer() {
	suite.Equal(19,Number(10, 20)  )
}


func (suite *Suite) TestNumerify() {
	suite.Equal("416",Numerify("###") )
}

func (suite *Suite) TestFixedNumberString() {
	suite.Equal("5108",FixedNumberString(4) )
}
