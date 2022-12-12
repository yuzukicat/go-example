package main

import (
	"fmt"
	"regexp"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AnimalType struct {
	food       string
	locomotion string
	noise      string
}

var animalMap = make(map[string]*AnimalType)

func (animalType *AnimalType) InitAnimal(noa string, toa string) {
	switch toa {
	case "cow":
		cow := AnimalType{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		}
		animalMap[noa] = &cow
		fmt.Println("Created it!")
	case "bird":
		bird := AnimalType{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		}
		animalMap[noa] = &bird
		fmt.Println("Created it!")
	case "snake":
		snake := AnimalType{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		}
		animalMap[noa] = &snake
		fmt.Println("Created it!")
	default:
		fmt.Println("animal do not have the type you want to create")
	}
}

func (animalType *AnimalType) QueryAnimal(noa string, aoa string) {
	var animal Animal
	animalType = animalMap[noa]
	if animalType != nil {
		animal = animalType
		switch aoa {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("animal do not have the action you want to query")
		}
	} else {
		fmt.Println("you have not created the animal you want to query")
	}
}

func (animalType *AnimalType) Eat() {
	fmt.Println(animalType.food)
}

func (animalType *AnimalType) Move() {
	fmt.Println(animalType.locomotion)
}

func (animalType *AnimalType) Speak() {
	fmt.Println(animalType.noise)
}

func (animalType *AnimalType) UserInputHandler(userInput string) {
	r := regexp.MustCompile(`(^newanimal|^query)[^\>](.+)[^\>](cow$|bird$|snake$|eat$|move$|speak$)`)
	res := r.FindAllStringSubmatch(userInput, -1)
	if res != nil {
		req := res[0][1]
		noa := res[0][2]
		toaoraoa := res[0][3]
		// fmt.Println(req)
		// fmt.Println(noa)
		// fmt.Println(toaoraoa)
		switch req {
		case "newanimal":
			animalType.InitAnimal(noa, toaoraoa)
		case "query":
			animalType.QueryAnimal(noa, toaoraoa)
		default:
			fmt.Println("the first string should be newanimal or query")
		}
	} else {
		fmt.Println("please check your input")
	}
}

func main() {
	for i := 0; i > -1; i++ {
		var animalType AnimalType
		var req string
		fmt.Scanln(&req)
		animalType.UserInputHandler(req)
	}
}
