package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

type Name struct {
	fname string
	lname string
}

func BinaryToString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

func (n Name) read() {
	var filename string
	_, err := fmt.Scan(&filename)
	if err != nil {
		fmt.Println(err)
	} else {
		datBinary, e := ioutil.ReadFile(filename)
		if e != nil {
			fmt.Println(e)
		} else {
			datString := BinaryToString(datBinary)
			r := regexp.MustCompile(`(?P<Fname>[A-Z,a-z]+)(\s{1})(?P<Lname>[A-Z,a-z]+)(\n{1})`)
			res := r.FindAllStringSubmatch(datString, -1)
			nameSlice := make([]Name, 0)
			for i := range res {
				n.fname = res[i][1][0:20]
				n.lname = res[i][3][0:20]
				nameSlice = append(nameSlice, n)
			}
			fmt.Print(nameSlice)
		}
	}
}

func main() {
	var n Name
	n.read()
}
