package findian

import (
	"fmt"
	"regexp"
)

func Findian() {
	var userinput string
	fmt.Scanln(&userinput)
	match, _ := regexp.MatchString("^[i,I].*[a,A].*[n,N]$", userinput)
	if match {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
