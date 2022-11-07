package main

import (
	"encoding/json"
	"fmt"
)

func Makejson() {
	var name string
	var address string
	_, err := fmt.Scanln(&name)
	_, err2 := fmt.Scanln(&address)
	if err != nil && err2 != nil {
		fmt.Println(err, err2)
	} else {
		userMap := map[string]string{
			"name":    name,
			"address": address,
		}
		// fmt.Println(userMap)
		userJsonObject, err3 := json.Marshal(userMap)
		if err3 != nil {
			fmt.Println(err3)
		} else {
			fmt.Printf("%s\n", userJsonObject)
		}
	}
}

func main() {
	Makejson()
}
