package main

import (
	"fmt"
)

func BubbleSort() {
	intSlice := make([]int, 10)
	for index := 0; index < cap(intSlice); index++ {
		var intItem int
		_, err := fmt.Scanln(&intItem)
		if err != nil {
			fmt.Println(err)
		} else {
			intSlice[index] = intItem
		}
	}
	Swap(intSlice)
}

func Swap(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	fmt.Println(slice)
	return slice
}

func main() {
	BubbleSort()
}
