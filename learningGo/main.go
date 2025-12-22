package main

import (
	"fmt"
	"learning/chapter2"
)

func main() {
	playground()
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
