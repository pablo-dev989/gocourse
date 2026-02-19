package main

import "fmt"

func main() {

	// func <name>(parameters list) returnType {
	// code block
	// return value
	// }

	//sum := add(2, 2)
	// fmt.Println(add(400, 50))

	// greet := func() {
	// 	fmt.Println("Esta es una funcion anonima")
	// }

	// greet()

	// operation := add

	// result := operation(3, 5)

	// fmt.Println(result)

	// Passing a function as an Argument
	result := applyOperation(5, 3, add)
	fmt.Println("5 + 3: ", result)

	// Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2: ", multiplyBy2(6))

}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that return a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
