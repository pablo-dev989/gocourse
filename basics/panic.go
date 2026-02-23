package basics

import "fmt"

func main() {

	// panic(interface())

	// Example of the valid input
	process(10)
	// Example of the invalid input
	process(-3)
}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before panic")
		panic("input must de a non-negative number")
		//fmt.Println("After panic")
	}
	fmt.Println("Processing input:", input)
}
