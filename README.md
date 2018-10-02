[![Build Status](https://travis-ci.org/ZhaoLion/faker.svg?branch=master)](https://travis-ci.org/ZhaoLion/faker)

# faker

faker(inspired by ruby gem faker) is a useful tool for generating data. 

Demo:

```
package main

import (
	"fmt"

	"github.com/ZhaoLion/faker"
)

func main() {
	faker.SetLocale("zh-CN")
	fmt.Printf("RandomCellPhone: %s\n", faker.RandomCellPhone())
}
```