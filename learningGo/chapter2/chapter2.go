package chapter2

import "fmt"

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

func E3() {

}
