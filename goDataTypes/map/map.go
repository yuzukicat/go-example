package main

import (
	"fmt"
)

// Notice that we've added items to a map that was initialized.
// But if you try to do the same with a nil map, you'll get an error.
// func studentAgeNilMap() {
// 	var studentAgeNilMap map[string]int
// 	studentAgeNilMap["John"] = 28
// }

func main() {
	// Declare and initialize a map
	studentAgesMap := map[string]int{
		"John": 28,
		"Blob": 26,
	}
	fmt.Println(studentAgesMap)

	studentAgeMap := make(map[string]int)
	fmt.Println(studentAgeMap)
	studentAgeMap["John"] = 28
	studentAgeMap["Blob"] = 26
	fmt.Println(studentAgeMap)

	// Notice that we've added items to a map that was initialized.
	// But if you try to do the same with a nil map, you'll get an error.

	//
	studentAgeChristy := studentAgeMap["Christy"]
	fmt.Println(studentAgeChristy)
}
