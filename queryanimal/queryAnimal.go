package main

import (
	"fmt"
	"regexp"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (p *Animal) UserInput() {
	var userinput string
	fmt.Scanln(&userinput)
	r := regexp.MustCompile(`(^cow|^bird|^snake)+[^\>]*(eat$|move$|speak$)+`)
	res := r.FindAllStringSubmatch(userinput, -1)
	name := res[0][1]
	p.InitAnimal(name)
	action := res[0][2]
	p.AnimalAction(action)
}

func (p *Animal) InitAnimal(name string) {
	if name == "cow" {
		p.food = "grass"
		p.locomotion = "walk"
		p.noise = "moo"
	} else if name == "bird" {
		p.food = "worms"
		p.locomotion = "fly"
		p.noise = "peep"
	} else if name == "snake" {
		p.food = "mice"
		p.locomotion = "slither"
		p.noise = "hsss"
	}
}

func (p *Animal) AnimalAction(action string) {
	if action == "eat" {
		p.Eat()
	} else if action == "move" {
		p.Move()
	} else if action == "speak" {
		p.Speak()
	}
}

func (p *Animal) Eat() {
	fmt.Print(p.food)
}

func (p *Animal) Move() {
	fmt.Print(p.locomotion)
}

func (p *Animal) Speak() {
	fmt.Print(p.noise)
}

func main() {
	var animal Animal
	animal.UserInput()
}
