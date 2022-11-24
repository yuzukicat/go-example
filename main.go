package main

import (
	"fmt"
	// "go-example/data"
	// "go-example/findian"
	// "go-example/hello"
	// "go-example/trunc"
	// "go-example/slice"
	"math"
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
// Associating Methods with Data: M	// "go-example/hello"ethod has a receiver type that it is associated with

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

// Composite data types: arrays, slices, maps, structs
// Aggregate
// Arrays: Fixed-length series of elements of a choosen type
// Elements accessed using subscript notion, []
// Indices start at 0
// Elements initialized to zero value

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

// Passing array pointers
func fooArr(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}

// slice contain a pointer to the array
// Passing a slice copies the pointer
// Can use the pointer directly to modify the slice without dereference/reference
// when defining a slice, the size of the slice can not be defined
func fooSlice(sli []int) {
	sli[0] = sli[0] + 1
}

// Code = functions + data
// Features can be found quickly
// Where data is defined can be found quickly
// Debugging from functions (understandable) or data (traceable)
// Function Cohesion: Function should perform only one operation
// Less number of parameters
// Function Complexity: Function length(lines)
// Use Function Call Hierarchyv
// Control-flow Comlexity: Partitioning Conditions

// Function as First-Class
// Function Type
// Dynamically create (Define) Function
// Function can be passed as arguments (to another function) and returned as values
// Function can be stored in data structures
var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

func decFn(x int) int {
	return x - 1
}

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

// Anonymous Function

func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}

// Closure: Function + its Environment

// Variadic: Pass a slice to a variadic function
func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

// Deffered Functions: Functions can be Deffered until the surronding function complete
// Be used in clean up
// Arguements of a deffered call are evaluated immediately

func main() {

	evaluateI := 1
	evaluateI++
	defer fmt.Print(evaluateI)

	// // trunc assignment
	// trunc.Trunc()

	// // Call by reference(Copy)
	// v := MyInt(3)
	// fmt.Println(v.Double())

	// // Encapsulation: Controlling Access
	// // Can define public functions to allow acacess to hidden data
	// data.PrintX()
	// // No need to reference
	// var p data.Point
	// p.InitMe(3, 4)
	// p.Scale(2)
	// p.PrintMe()

	// foo(2, 3)
	// // variable assignment: a, b to get captured
	// // :=call the func, pass the Arguments (2, 3), (2, 3) is going to get bound to x inside the function foo when it gets executed.
	// // (2, 3) is get evaluated, and be replaced with (2, 5), when excuted, (2, 5) is assigned to (a, b)
	// // return value must have an assignment.
	// a, b := foo2(2, 3)
	// print(a, b)

	// x := 2
	// // pass a point to x, pass that to foor, foor have access to (a copy of the location in the memory where x is)
	// // so foor modifying the data at that location
	// // foor has the ability to alter variable x (even though x was not initially defined inside the scope of foo)
	// // x was defined in the scope of main
	// foor(&x)
	// fmt.Print(x)

	aArr := [3]int{1, 2, 3}
	fooArr(&aArr)
	fmt.Print(aArr)

	aSli := []int{1, 2, 3}
	fooSlice(aSli)
	fmt.Print(aSli)

	// hello.Hello()

	// findian.Findian()

	// var ar [5]int
	// ar[0] = 2
	// // An array pre-defined with values
	// // ...for size in array lteral infers size from number of initializers
	// ar2 := [...]int{1, 2, 3, 4, 5}
	// fmt.Println(ar2)
	// for ind, val := range ar2 {
	// 	fmt.Println(ind, val)
	// }

	// // Slices: A window on an underlying array
	// // Slices properties: Pointer, indicates the start of the slice
	// // Length
	// // Capacity, From start of slice to end of array
	// arr := [...]string{"a", "b", "c", "d", "e", "", "g"}
	// sli1 := arr[1:3]
	// sli2 := arr[2:5]
	// fmt.Println(len(sli1), cap(sli2))
	// // Accessing Slices
	// // Writing to a slice changes underlying array and
	// // Overlapping slices refer to the same array elements
	// fmt.Println(sli1[1])
	// fmt.Println(sli2[0])
	// // Slice Literals, can be used to initialize a slice
	// // Creates the underlying array and creates a slice to reference it
	// // Slice points to the start of the array, length is capacity
	// sli3 := []int{1, 2, 3}
	// fmt.Println(sli3)
	// // Make a slice
	// // 2-argument version: specify type and length/capcity
	// sli4 := make([]int, 0, 3)
	// fmt.Println(sli4)
	// sli5 := make([]int, 10, 15)
	// fmt.Println(sli5)
	// // Append. size of a slice can be increased by append()
	// // Inserts into underlying array
	// // Increases size of array if necessary
	// sli4 = append(sli4, 100)
	// fmt.Println(sli4)
	// // Hash table, contains key/value pairs
	// // Each vallue is associated with a unique key
	// // Hash function is used to compute the slot (return) for a key (argument)
	// // Advantages of hash tables: Constant-time vs linear-time, faster lookup than lists
	// // Arbitrary keys.
	// // Disadvantages. May have collisions: Two keys hash to the same slot
	// // Maps. Golang implementation of a hash table
	// // Use make() to create a map
	// var idMap map[string]int
	// idMap2 := make(map[string]int)
	// idMap3 := map[string]int{
	// 	"joe": 123,
	// }
	// fmt.Println(idMap)
	// fmt.Println(idMap2)
	// fmt.Println(idMap3["joe"])
	// // Accessing Maps. Referencing a value with [key]
	// // returns zero if key is not present
	// // Adding/change a key/value pair
	// idMap3["jane"] = 456
	// fmt.Println(idMap3)
	// // Deleting a key/value pair
	// delete(idMap3, "joe")
	// fmt.Println(idMap3)
	// // Map Functions
	// // Two-value assignment tests for existence of the key
	// id, p := idMap3["jane"] // p is boolean (presence of key)
	// fmt.Println(id, p)
	// // len() returns number of values
	// fmt.Println(len(idMap3))
	// // Iterrating through a map
	// // Use a for loop with the range keyword
	// // Two-value assignment with range
	// for key, val := range idMap3 {
	// 	fmt.Println(key, val)
	// }

	// Struct: Aggregate data type
	// Groups together other objects of arbitrary type
	// For organizational purposes, it is helps
	// type struct Person {
	// 	//field
	// 	name string
	// 	addr string
	// 	phone string
	// }
	// var p1 Person
	// p1.name = "joe"
	// x = p1.name
	// fmt.Println(x)

	// slice.SliceScanln()

	funcVar = incFn
	fmt.Print(funcVar(1))

	fmt.Print(applyIt(incFn, 1))
	fmt.Print(applyIt(decFn, 1))

	fmt.Print(getMax(1, 3, 4, 6))

	vslice := []int{1, 3, 6, 4}

	fmt.Print(getMax(vslice...))

}
