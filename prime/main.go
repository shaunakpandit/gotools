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
		fmt.Println("Expecting a number. Value given: ", i)
	}

	abs(&num)

	fmt.Println(num, " is prime: ", isPrime(num, num/2))
}

func abs(n *int) int {
	if *n > 0 {
		return *n
	} else if *n == 0 {
		return 0
	} else {
		return -1 * *n
	}
}

func isPrime(num, div int) bool {
	if num == 0 || num == 1 {
		return false
	}
	if div == 1 {
		return true
	}
	if num%div == 0 {
		return false
	}
	return isPrime(num, div-1)
}
