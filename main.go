package main

import "fmt"

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
