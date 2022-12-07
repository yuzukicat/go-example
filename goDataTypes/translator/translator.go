package main

import "fmt"

func RomanToArabic(input string) int {
	romanArabicMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	// for the porpose of +-, the cap should be len+1
	arabicValues := make([]int, len(input)+1)

	// get info from roman
	for index, roman := range input {
		if arabic, presence := romanArabicMap[roman]; presence {
			arabicValues[index] = arabic
		} else {
			fmt.Printf("Error: The input have a bad digit:%s\t%c\n", input, roman)
			return 0
		}
	}

	arabicSum := 0

	// the len should be the input
	for i := 0; i < len(input); i++ {
		if arabicValues[i] < arabicValues[i+1] {
			arabicValues[i] = -arabicValues[i]
		}
		arabicSum += arabicValues[i]
	}

	return arabicSum
}

func main() {
	fmt.Println(RomanToArabic("MCM"))
}
