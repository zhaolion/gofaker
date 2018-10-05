package main

import (
	"fmt"

	"github.com/zhaolion/faker"
)

func main() {
	faker.SetLocale("zh-CN")
	fmt.Printf("RandomCellPhone: %s\n", faker.RandomCellPhone())
}
