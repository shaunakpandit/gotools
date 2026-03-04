package main

import (
	"fmt"
	"math/rand"
	"sortingAlgos/bubblesort"
	"time"
)

func main() {
	randSlice := BuildSlice(20)
	fmt.Println("unsorted:", randSlice)
	output := bubble(randSlice)
	fmt.Println("sorted:", output)
}

func bubble(unsorted []int) []int {
	start := time.Now()
	sorted := bubblesort.BubbleSort(unsorted)
	end := time.Now()
	fmt.Println("bubbleSort took:", end.Sub(start).Nanoseconds(), "nano")
	return sorted
}

func BuildSlice(size int) []int {
	randSlice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		randSlice = append(randSlice, rand.Intn(size*10))
	}
	return randSlice
}
