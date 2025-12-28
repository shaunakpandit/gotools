package main

import (
	"fmt"
	"learning/chapter2"
	"learning/chapter3"
	"learning/chapter4"
)

func main() {
	chapter4.E3()
}

func chap2() {
	chapter2.All()
}

func chap3() {
	chapter3.All()
}

func chap4() {
	chapter4.All()
}

func playground() {
	m := map[string]int{
		"hello": 5,
		"world": 3,
	}

	for k, v := range m {
		fmt.Println(m[k])
		fmt.Println(v)
	}
}
