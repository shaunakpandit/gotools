package main

import (
	"fmt"
	"learning/chapter2"
	"learning/chapter3"
)

func main() {
	chapter3.All()
}

func chap2() {
	chapter2.All()
}

func playground() {
	m := map[string]int{
		"hello": 5,
		"world": 3,
	}

	v, ok := m["no"]

	fmt.Println(v, ok)
}
