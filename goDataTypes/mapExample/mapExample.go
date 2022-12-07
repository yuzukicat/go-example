package mapExample

import (
	"fmt"
)

// Notice that we've added items to a map that was initialized.
// But if you try to do the same with a nil map, you'll get an error.
// func studentAgeNilMap() {
// 	var studentAgeNilMap map[string]int
// 	studentAgeNilMap["John"] = 28
// }

func MapExample() {
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
	studentAgeChristy, exist := studentAgeMap["Christy"]
	if exist {
		fmt.Println(studentAgeChristy)
	} else {
		fmt.Println("Not Found")
	}

	// Remove items
	delete(studentAgeMap, "Blob")
	fmt.Println(studentAgeMap)

	// if you try to delete an item that doesn't exist, Go won't panic.

	// Loop in a map
	studentAgeMap["Blob"] = 26
	studentAgeMap["Christy"] = 20
	for name, age := range studentAgeMap {
		fmt.Printf("%s\t%d\n", name, age)
	}

}
