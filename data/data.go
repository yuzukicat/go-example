package data

import "fmt"

var x int = 1

func PrintX() {

	fmt.Println(x)

}

// Good programming practice:
// All methods for a type have pointer receivers
type Point struct {
	x float64
	y float64
}

// Pointer Receiver: Receiver can be a pointer to a type(call by value)
func (p *Point) InitMe(xn, yn float64) {
	p.x = xn
	p.y = yn
}

// No need to Dereference
// Point is referenced as p, not *p
// Dereferencing is automatic with .operator
func (p *Point) Scale(v float64) {
	p.x = p.x * v
	p.y = p.y * v
}

func (p *Point) PrintMe() {
	fmt.Println(p.x, p.y)
}
