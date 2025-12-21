package main

import (
	"learning/chapter2"
	"learning/chapter3"
)

func main() {
	playground()
}

func chap2() {
	chapter2.All()
}

func playground() {
	var x []int
	for len(x) < 10 {
		x = append(x, len(x)+1)
		chapter3.PrintSlice(x)
	}
}
