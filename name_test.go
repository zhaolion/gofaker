package faker

import "fmt"

func (suite *Suite) TestName() {
	var tests = []struct {
		locale string
		expect string
	}{
		{
			locale: "en",
			expect: "Truman Wintheiser III",
		},
	}

	for i, test := range tests {
		bundle.SetLocale(test.locale)
		suite.Equal(test.expect, Name(), fmt.Sprintf("case: %d", i))
	}
}
