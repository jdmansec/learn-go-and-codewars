# Go Variables and Types Instruction Guide

This document serves as an instruction manual for a Go program that demonstrates various aspects of working with variables. The program covers multiple concepts related to variable declaration, manipulation, and usage in Go.

## Table of Contents
1. Declaring Variables
2. Type Inference
3. Constants
4. Updating Variables
5. Printing Variables
6. Variable Operations
7. Short Declarations
8. Types of Variables
9. Zero Values

---

## 1. Declaring Variables

Go is a statically typed language, meaning the type of each variable is known at compile time. There are several ways to declare variables in Go.

#### a. Declaring a Single Variable with an Initial Value
*Here, we declare a string variable called `name` and initialize it with the value `"jdmansec"`. This demonstrates how to explicitly specify the type during variable declaration.*

```go
var name string = "jdmansec"
fmt.Printf("Variable 'name' has a value of %#v of type %T \n\n", name, name)
```
*Output:*
```
Variable 'name' has a value of "jdmansec" of type string
```

#### b. Declaring Multiple Variables of the Same Type
*Multiple variables (x, y, and z) can be declared together and initialized with values 1, 2, and 3, respectively.*

```go
var x, y, z int = 1, 2, 3
fmt.Printf("x = %#v \ny = %#v \nz = %#v \n\n", x, y, z)
```
*Output:*
```
x = 1
y = 2
z = 3
```

#### c. Declaring Variables Without Initialization
*Variables declared without an explicit initial value will be assigned the default "zero value" for their type.*

```go
var a int    // a is 0
var b string // b is ""
var c bool   // c is false

fmt.Printf("Variable 'a' has a value of %#v of type %T \n", a, a)
fmt.Printf("Variable 'b' has a value of %#v of type %T \n", b, b)
fmt.Printf("Variable 'c' has a value of %#v of type %T \n", c, c)
```
*Output:*
```
Variable 'a' has a value of 0 of type int
Variable 'b' has a value of "" of type string
Variable 'c' has a value of false of type bool
```

## 2. Type Inference

Go can infer the type of a variable from its initial value when using the shorthand declaration.

#### a. Using Inferred Types
*With the := operator, you can declare and initialize a variable without explicitly specifying its type.*

```go
inferredVariable := 42          // inferredVariable is an int
anotherVariable := "Hello, Jim" // anotherVariable is a string

fmt.Printf("%v is of type %T \n", anotherVariable, anotherVariable)
fmt.Printf("%v is of type %T \n", inferredVariable, inferredVariable)
```
*Output:*
```
Hello, Jim is of type string
42 is of type int
```

## 3. Constants

Constants are immutable values that are determined at compile time.

#### a. Declaring a Constant
*Here's how to declare a constant of type float64:*

```go
const pi float64 = 3.14159
fmt.Printf("%v is a constant (variable) of type %T \n", pi, pi)
```
*Output:*
```
3.14159 is a constant (variable) of type float64
```

## 4. Updating Variables

Variables declared with `var` can be modified throughout the program's execution.

#### a. Reassigning Values
*Variables can be updated with new values of the same type:*

```go
name = "Learning Go"
x = 10

fmt.Printf("%v is now assigned to 'name', and %#v is now assigned to 'x'\n", name, x)
```
*Output:*
```
Learning Go is now assigned to 'name', and 10 is now assigned to 'x'
```

## 5. Printing Variables

Go provides various functions in the `fmt` package for output formatting.

#### a. Using fmt.Println and fmt.Printf
*Different ways to print variables:*

```go
fmt.Println("Name:", name)
fmt.Println("x:", x, "y:", y, "z:", z)
fmt.Println("Uninitialized a:", a)
fmt.Println("Inferred Variable:", inferredVariable)
fmt.Println("Constant Pi:", pi)
```
*Output:*
```
Name: Learning Go
x: 10 y: 2 z: 3
Uninitialized a: 0
Inferred Variable: 42
Constant Pi: 3.14159
```

## 6. Variable Operations

Variables can be used in arithmetic and other operations.

#### a. Performing Arithmetic Operations
*Example of adding variables:*

```go
sum := x + y + z
fmt.Println("Sum:", sum)
```
*Output:*
```
Sum: 15
```

## 7. Short Declarations

Inside functions, Go provides a shorthand syntax for variable declaration and initialization.

#### a. Using the Short Declaration Operator
*The := operator combines declaration and initialization:*

```go
message := "Go makes it easy to declare variables"
fmt.Println(message)
```
*Output:*
```
Go makes it easy to declare variables
```

## 8. Types of Variables

Go supports various data types for different purposes.

#### a. Declaring Variables with Different Data Types
*Examples of different variable types:*

```go
var isGoFun bool = true
var temperature float64 = 36.6
var bigNumber int64 = 1 << 62

fmt.Println("Is Go Fun?", isGoFun)
fmt.Println("Temperature:", temperature)
fmt.Println("Big Number:", bigNumber)
```
*Output:*
```
Is Go Fun? true
Temperature: 36.6
Big Number: 4611686018427387904
```

## 9. Zero Values

Go automatically assigns default values to variables that aren't explicitly initialized.

#### a. Demonstrating Zero Values
*Default values for different types:*

```go
var zeroInt int
var zeroFloat float64
var zeroBool bool
var zeroString string

fmt.Println("Zero Int:", zeroInt)         // Prints: 0
fmt.Println("Zero Float:", zeroFloat)     // Prints: 0.0
fmt.Println("Zero Bool:", zeroBool)       // Prints: false
fmt.Println("Zero String:", zeroString)   // Prints: ""
```
*Output:*
```
Zero Int: 0
Zero Float: 0
Zero Bool: false
Zero String: 
```