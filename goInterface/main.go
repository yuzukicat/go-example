package main

import "fmt"

// Polymorphism

// Interfaces: Set of method signatures(name, paramaters, return values)
// Type satisfies an interfaces(can own additional methods)
// Interface: Dynamic Type (Concrete Type); Dynamic Value: Value of the dynamic (Pair)
// Interface can be treated like other values: Assigned to variables, passed, returned
// Interface not associated with data

type shape2D interface {
	Area() float64
	Perimeter() float64
}

func FitInYard(s shape2D) bool {
	if s.Area() > 100 &&
		s.Perimeter() > 100 {
		return true
	}
	return false
}

type Speaker interface{ Speak() }

type Dog struct {
	name string
}

// func (d Dog) Speak() {
func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("noise")
	} else {
		fmt.Println(d.name)
	}
}

func main() {
	var speaker Speaker
	dog := Dog{
		"Brian",
	}
	// 	speaker = dog
	speaker = &dog
	speaker.Speak()

	// the dynamic type for speaker is dog and dynamic value for speaker is name
	// Interface can have Nnil dynamic value
	// Nil dynamic value and valid dynamic type
	// When using pointer, speaker is not an actual dog but a pointer to a dog
	// Which means speaker have dynamic type pointer type and nil value (has no concreate value yet)
	// Compiler can figure out it is the dog's speak
	var dogNil *Dog
	speaker = dogNil
	speaker.Speak()

	// Nil dynamic type
	// Can not call a method, runtime error
	// Because interface does not give the implementation of the method
	// var speakerNil Speaker
}
