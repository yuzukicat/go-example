package main

import (
	"fmt"
	"go-example/data"
	"go-example/trunc"
)

// Advantages of go
// Code runs fast
// Garbage collection (Memory management)
// Simpler objects
// Concurrency primitives (Parallelism)
// Goruntines represent concurrent tasks
// Channels are used to communicate between tasks
// Select enables task synchronization

// Compilation: Translate instructions once before running the code.
// Interpreation: Translate instructions while code is executed.

// Object-Oriented Programming: Create Types
// Organize code through encapsulation
// Group together data and functions
// Ex. ints have data (number) and functions (+,-,*,etc)
// Go uses structs (data, does not use class/instance) with associated methods (functions)
// Struct is simplified implementation of classes (No inheritance, no constructors, no generics)

// Go workspace src (code files)/pkg (packages)/bin (executables)

// keyword(var) name type
// var x int
// var y float32
// var z string

// Type Declarations (alias)
// type celsius float64

// var temp celsius

// Innitializing Variables
// Initialize in the declaration
// var m int = 100

// var n = 100

// Uninitialized zero value

// Short Variable Declarations (Can only inside func)

// oq := 100

// Data types
// Pointers & Address * Data
// new() returns a pointer to obj

// Variable Scope
// Blocks
// Stack(auto deallocating after function complete)
// Heap is persistant (determined by garbage collection during compile time)

//printf
// %s is conversion character for an argument (be substituted in %s)

// String, Read-one, 0x41=A is a code point in unicode (rune in go)
// String, sequence of arbitrary bytes
// String literal, x := "Hi there"

// String Manipulation

// Type conversions

// Constants

type Grades int

const (
	A Grades = iota
	B
	C
)

// for <init(option)>; <cond>; <update(option)> {
// 		<stmts>
// 		<update(option)>
// }

// Switch / Case, x is a tag which is a variable to be checked

// Tagless Switch. (excute first one, else if like)

// Break continue

// Scan

// Class(template): Collection of data fields(not data) and functions
// Associating Methods with Data: Method has a receiver type that it is associated with

// Object: Instance of class, contains real data
// Encapssulation: Data can be accessed only using methods
// Method has a receiver type that is associated with
type MyInt int

func (mi MyInt) Double() int {
	return int(mi * 2)
}

// struct types compose data fields
// type Point struct {
// 	x float64
// 	y float64
// }

// func (p Point) DistToOrigin() {
// 	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
// 	return math.Sqrt(t)
// }

// Why use functions?
// Reusability
// Abstraction (Hiddle details, improve understandability)
func PrintHello() {
	fmt.Printf("Hello, world.")
}

// Parameters are listed in parenthesis after function name
// Arguments are supplied in the call
func foo(x int, y int) {
	fmt.Print(x + y)
}

// List arguments of same type
// Type of return value after parameters in declartion
// Multiple Return Values. Multiple value types must be listed in the declaration
func foo2(x, y int) (int, int) {
	return x, x + y
}

// Tradeoffs of call by value
// Advantage: Data Encapsulation
// Disadvantage: Copying time

// Call by Reference (pass a pointer as an argument)
func foor(y *int) {
	// y is a pointer to an integer (it knows where that value is in the memory)
	// so y can directly go into that location of memory and alter it
	*y = *y + 1
}

func main() {

	// trunc assignment
	trunc.Trunc()

	// Call by reference(Copy)
	v := MyInt(3)
	fmt.Println(v.Double())

	// Encapsulation: Controlling Access
	// Can define public functions to allow acacess to hidden data
	data.PrintX()
	// No need to reference
	var p data.Point
	p.InitMe(3, 4)
	p.Scale(2)
	p.PrintMe()

	foo(2, 3)
	// variable assignment: a, b to get captured
	// :=call the func, pass the Arguments (2, 3), (2, 3) is going to get bound to x inside the function foo when it gets executed.
	// (2, 3) is get evaluated, and be replaced with (2, 5), when excuted, (2, 5) is assigned to (a, b)
	// return value must have an assignment.
	a, b := foo2(2, 3)
	print(a, b)

	x := 2
	// pass a point to x, pass that to foor, foor have access to (a copy of the location in the memory where x is)
	// so foor modifying the data at that location
	// foor has the ability to alter variable x (even though x was not initially defined inside the scope of foo)
	// x was defined in the scope of main
	foor(&x)
	fmt.Print(x)
}
