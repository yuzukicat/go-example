package sliceExample

import (
	"fmt"
)

func SliceExample() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(len(months))
	fmt.Println(cap(months))

	// The element at position p in the underlying array is found at the location array[i+1].
	// Notice that this element isn't necessarily the underlying array's last element, array[len(array)-1].
	quarter2 := months[3:6]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	quarter2Extend := quarter2[:4]
	fmt.Println(quarter2Extend, len(quarter2Extend), cap(quarter2Extend))

	// When a slice doesn't have enough capacity to hold more elements, Go doubles its capacity.
	// It creates a new underlying array with the new capacity.
	// At some point, a slice might have way more capacity than it needs, and you'll be wasting memory.
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Println(i, cap(numSlice), numSlice)
	}

	letters := []string{"A", "B", "C", "D", "E"}
	removePosition := 2
	fmt.Println("BeforeRemove", letters, "ToBeRemoved", letters[removePosition])
	if removePosition < len(letters) {
		// It replaces the element to be removed with the next element in the slice,
		// or none if you're removing the last element.
		letters = append(letters[:removePosition], letters[removePosition+1:]...)
		fmt.Println("AfterRemove", letters)
	}

	// a slice points to an array, and every change you make in a slice affects the underlying array.
	lettersSlice1 := letters[0:4]
	lettersSclice2 := letters[1:2]
	lettersSclice2[0] = "Z"
	fmt.Println(lettersSlice1)

	lettersSclice3 := make([]string, 4)
	copy(lettersSclice3, lettersSlice1)
	lettersSclice2[0] = "W"
	fmt.Println(lettersSclice3)
}
