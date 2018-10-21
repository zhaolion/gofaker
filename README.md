[![Build Status](https://travis-ci.org/zhaolion/gofaker.svg?branch=master)](https://travis-ci.org/zhaolion/gofaker)[![GoDoc](https://godoc.org/github.com/zhaolion/gofaker?status.svg)](https://godoc.org/github.com/zhaolion/gofaker)
# gofaker

gofaker(inspired by ruby gem faker) is a useful tool for generating data.

**go version: 1.9+**

## Usages
- [Alphanumeric](doc/alphanumeric.md)
- [Name](doc/name.md)
- [PhoneNumber](doc/phone_number.md)


## Installation
`go get github.com/zhaolion/gofaker`

## Demo:

```
package main

import (
	"fmt"

	"github.com/zhaolion/gofaker"
)

func main() {
	faker.SetLocale("zh-CN")
	fmt.Println(gofaker.RandomCellPhone())
}
```

## How to contribute

```
curl https://glide.sh/get | sh
make dep
make test
```

Please PR, Thanks.
