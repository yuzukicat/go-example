package main

import "fmt"

func Fibonacci(fibonacciLength int) []int {
	if fibonacciLength < 2 {
		return make([]int, 0)
	}

	numSlice := make([]int, fibonacciLength)
	numSlice[0], numSlice[1] = 1, 1

	for i := 2; i < fibonacciLength; i++ {
		numSlice[i] = numSlice[i-1] + numSlice[i-2]
	}
	return numSlice
}

func main() {
	var fibonacciLength int
	fmt.Scanln(&fibonacciLength)
	fmt.Println(Fibonacci(fibonacciLength))
}
