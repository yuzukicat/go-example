package animal

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
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

func (p *Animal) Eat() {
	fmt.Print(p.food)
}

func (p *Animal) Move() {
	fmt.Print(p.locomotion)
}

func (p *Animal) Speak() {
	fmt.Print(p.noise)
}
