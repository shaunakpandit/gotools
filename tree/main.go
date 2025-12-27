package main

import (
	"fmt"
)

func main() {
	baseSize := 15
	// fmt.Println("*")
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
