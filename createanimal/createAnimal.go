package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AnimalAttr struct {
	food       string
	locomotion string
	noise      string
}

func (animalAttr *AnimalAttr) InitAnimal(name string) {
	animalMap := make(map[string]AnimalAttr)
	fmt.Scanln(&name)
	switch name {
	case "cow":
		animalMap["cow"] = AnimalAttr{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		}
		fmt.Println("Created it!")
	case "bird":
		animalMap["bird"] = AnimalAttr{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		}
		fmt.Println("Created it!")
	case "snake":
		animalMap["snake"] = AnimalAttr{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		}
		fmt.Println("Created it!")
	}
}

func (animalAttr *AnimalAttr) QueryAnimal(name string, action string) {

}

func (animalAttr *AnimalAttr) Eat() {
	fmt.Println(animalAttr.food)
}

func (animalAttr *AnimalAttr) Move() {
	fmt.Println(animalAttr.locomotion)
}

func (animalAttr *AnimalAttr) Speak() {
	fmt.Println(animalAttr.noise)
}
