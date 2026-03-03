package main

import (
	"fmt"
	"math/rand"
	"sortingAlgos/bubblesort"
)

func main() {
	randSlice := BuildSlice(10)
	output := bubblesort.BubbleSort(randSlice)
	fmt.Print(output)
}

func BuildSlice(size int) []int {
	randSlice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		randSlice = append(randSlice, rand.Intn(100))
	}
	return randSlice
}
