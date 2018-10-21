package main

import (
	"fmt"

	"github.com/zhaolion/gofaker"
)

func main() {
	gofaker.SetLocale("zh-CN")
	fmt.Printf("RandomCellPhone: %s\n", gofaker.RandomCellPhone())
}
