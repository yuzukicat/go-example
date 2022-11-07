package slice

import (
	"fmt"
	"sort"
)

func SliceScanln() {
	emptyIntegerSlice := make([]int, 0)
	for round := -1; round < cap(emptyIntegerSlice); round++ {
		var userinput int
		_, err := fmt.Scanln(&userinput)
		if err != nil {
			break
		} else {
			emptyIntegerSlice = append(emptyIntegerSlice, userinput)
			sort.Sort(sort.IntSlice(emptyIntegerSlice))
			fmt.Println(emptyIntegerSlice)
		}
	}
}
