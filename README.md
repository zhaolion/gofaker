[![Build Status](https://travis-ci.org/zhaolion/faker.svg?branch=master)](https://travis-ci.org/ZhaoLion/faker)[![GoDoc](https://godoc.org/github.com/zhaolion/faker?status.svg)](https://godoc.org/github.com/zhaolion/faker)
# faker

faker(inspired by ruby gem faker) is a useful tool for generating data. 

## Usages
- [Alphanumeric](doc/alphanumeric.md)
- [PhoneNumber](doc/phone_number.md)


## Installation
`go get github.com/ZhaoLion/faker`

## Demo:

```
package main

import (
	"fmt"

	"github.com/ZhaoLion/faker"
)

func main() {
	faker.SetLocale("zh-CN")
	fmt.Println(faker.RandomCellPhone())
}
```

## How to contribute

```
curl https://glide.sh/get | sh
make dep
make test
```

Please PR, Thanks.
