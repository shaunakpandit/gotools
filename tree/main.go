package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i := os.Args[1]
	num, err := strconv.Atoi(i)

	if err != nil {
		message := "Expecting a number. Value given: " + i
		panic(message)
	}
	if num%2 == 0 {
		message := "Expecting an odd number. Value given: " + i
		panic(message)
	}

	baseSize := num
	for i := 0; i <= baseSize; i += 2 {
		layerSpaces := (baseSize - i) / 2
		layerStars := baseSize - layerSpaces*2

		for range layerSpaces {
			fmt.Print(" ")
		}

		for range layerStars {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for range baseSize / 2 {
		fmt.Print(" ")
	}
	fmt.Println("*")
}
