package main

import (
	"fmt"
	"math/rand"
	"sortingAlgos/bubblesort"
)

func main() {
	randSlice := BuildSlice(20)
	fmt.Println("unsorted:", randSlice)
	output := bubblesort.BubbleSort(randSlice)
	fmt.Println("sorted:", output)
}

func BuildSlice(size int) []int {
	randSlice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		randSlice = append(randSlice, rand.Intn(size*10))
	}
	return randSlice
}
