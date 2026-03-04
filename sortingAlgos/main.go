package main

import (
	"fmt"
	"math/rand"
	"sortingAlgos/bubblesort"
	"time"
)

// Sorting Algorithms (Easiest → Hardest)
//
// 1. Bubble Sort (done)
// 2. Selection Sort
// 3. Insertion Sort
// 4. Merge Sort
// 5. Quick Sort
// 6. Heap Sort
// 7. Shell Sort
// 8. Counting Sort
// 9. Radix Sort
// 10. Bucket Sort
// 11. Tim Sort
// 12. Intro Sort

func main() {
	bubbleTimed(1000, 100, false)
}

func bubbleTimed(runs int, size int, print bool) {
	var totalTime int64 = 0
	for range runs {
		randSlice := BuildSlice(size)
		output, time := bubble(randSlice)
		totalTime += time
		if print {
			fmt.Println("unsorted:", randSlice)
			fmt.Println("sorted:", output)
			fmt.Println("bubbleSort took:", time, "nano")
		}
	}
	fmt.Println("average time to sort", size, "items, over", runs, "runs:", totalTime/int64(runs), "nanos")
}

func bubble(unsorted []int) ([]int, int64) {
	start := time.Now()
	sorted := bubblesort.BubbleSort(unsorted)
	end := time.Now()
	duration := end.Sub(start).Nanoseconds()
	return sorted, duration
}

func BuildSlice(size int) []int {
	randSlice := make([]int, 0, size)
	for range size {
		randSlice = append(randSlice, rand.Intn(size*10))
	}
	return randSlice
}
