package arrayExample

import (
	"fmt"
)

func ArrayExample() {
	// Declare arrays
	var arrVar [3]int
	arrVar[1] = 10
	fmt.Println(arrVar)
	fmt.Println(len(arrVar))

	// Initialize arrays
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("cities", cities)

	// Ellipsis in arrays
	// There's no empty string at the end.
	// The array length was determined by the strings you put when you initialized it.
	// You're not reserving memory you don't know if you'll end up needing.
	numArr := [...]int{1, 2, 3}
	fmt.Print(numArr)

	lastPositionArray := [...]int{99: -1}
	fmt.Println("Length:", len(lastPositionArray))
	fmt.Println("Last Position:", lastPositionArray[99])

	// Multidimensional arrays
	var twoDmArr [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoDmArr[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("ROW", i, twoDmArr[i])
	}
	fmt.Println("\nAll at once:", twoDmArr)
}
