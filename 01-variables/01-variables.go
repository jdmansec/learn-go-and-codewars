package main

import (
	"fmt"
)

func main() {
	// 1. Declaring Variables
	// Go is a statically typed language, so the type of a variable is known at compile time.
	fmt.Println()
	fmt.Println("1. Declaring Variables")
	fmt.Println()

	// Declaring a single variable with an initial value
	var name string = "Jim Dearman"

	fmt.Printf("Variable 'name' has a value of %#v of type %T \n\n", name, name)

	// Declaring multiple variables of the same type
	var x, y, z int = 1, 2, 3

	fmt.Printf("x = %#v \ny = %#v \nz = %#v \n\n", x, y, z)

	// Declaring variables without initializing them
	// These will take the "zero value" for their type
	var a int    // a is 0
	var b string // b is ""
	var c bool   // c is false

	fmt.Printf("Variable 'a' has a value of %#v of type %T \n", a, a)
	fmt.Printf("Variable 'b' has a value of %#v of type %T \n", b, b)
	fmt.Printf("Variable 'c' has a value of %#v of type %T \n", c, c)

	// 2. Type Inference
	fmt.Println()
	fmt.Println("2. Type Inference")
	fmt.Println()
	// You can omit the type if the variable is initialized,
	// Go will infer the type from the initial value.
	inferredVariable := 42          // inferredVariable is an int
	anotherVariable := "Hello, Jim" // anotherVariable is a string

	fmt.Printf("%v is of type %T \n", anotherVariable, anotherVariable)
	fmt.Printf("%v is of type %T \n", inferredVariable, inferredVariable)

	// 3. Constants
	fmt.Println()
	fmt.Println("3. Constants")
	fmt.Println()
	// Constants are variables whose values cannot be changed.
	const pi float64 = 3.14159

	fmt.Printf("%v is a constant (varaible) of type %T \n", pi, pi)

	// 4. Updating Variables
	fmt.Println()
	fmt.Println("4. Updating Variables")
	fmt.Println()
	// Variables declared with var can be updated later.
	name = "Learning Go"
	x = 10

	fmt.Printf("%v is now assigned to 'name', and %#v is now assigned to 'x'\n", name, x)

	// 5. Printing Variables
	fmt.Println()
	fmt.Println("5. Printing Variables")
	fmt.Println()
	// Use fmt.Println to print variables to the console
	fmt.Println("Name:", name)
	fmt.Println("x:", x, "y:", y, "z:", z)
	fmt.Println("Uninitialized a:", a)
	fmt.Println("Inferred Variable:", inferredVariable)
	fmt.Println("Constant Pi:", pi)

	// 6. Variable Operations
	fmt.Println()
	fmt.Println("6. Variable Operations")
	fmt.Println()
	sum := x + y + z

	fmt.Println("Sum:", sum)

	// 7. Short Declarations (Inside Functions Only)
	fmt.Println()
	fmt.Println("7. Short Declarations (Inside Functions Only)")
	fmt.Println()
	// The := syntax is a shorthand for declaring and initializing a variable.
	message := "Go makes it easy to declare variables"
	fmt.Println(message)

	// 8. Types of Variables
	fmt.Println()
	fmt.Println("8. Types of Variables")
	fmt.Println()
	var isGoFun bool = true
	var temperature float64 = 36.6
	var bigNumber int64 = 1 << 62

	fmt.Println("Is Go Fun?", isGoFun)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Big Number:", bigNumber)

	// 9. Zero Values
	fmt.Println()
	fmt.Println("9. Zero Values")
	fmt.Println()
	// If you declare a variable without assigning it an initial value,
	// Go will assign it a zero value.
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string

	fmt.Println("Zero Int:", zeroInt)
	fmt.Println("Zero Float:", zeroFloat)
	fmt.Println("Zero Bool:", zeroBool)
	fmt.Println("Zero String:", zeroString)
}
