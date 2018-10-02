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