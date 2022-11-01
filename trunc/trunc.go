package trunc

import (
	"fmt"
	"math"
	"strconv"
)

func Trunc() {
	// Println function is used to
	// display output in the next line
	fmt.Println("Enter A Floating Point Number: ")
	// var then variable name then variable type
	var strInput string
	// Taking input from user
	fmt.Scanln(&strInput)
	fmt.Printf("You Entered: %s", strInput)
	fmt.Println("removing...")
	if floatingInput, err := strconv.ParseFloat(strInput, 64); err == nil {
		intInput := math.Trunc(floatingInput)
		fmt.Println(intInput)
	}
}
