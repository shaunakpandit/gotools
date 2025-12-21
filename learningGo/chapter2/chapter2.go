package chapter2

import (
	"fmt"
)

func E1() {
	var i int = 20
	var f float64 = float64(i)
	fmt.Println(i)
	fmt.Println(f)
}

func E2() {
	const value = 10
	var i int = value
	var f float32 = value
	fmt.Println(i)
	fmt.Println(f)
}

// literals will overflow to min value when exceeding an upper limit
func E3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)
}
