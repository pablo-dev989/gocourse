package main

import "fmt"

func main() {

	// func functionName(parameter1 type1, parameter2 type2,...) (returnType1, returnType2,...){
	////// code block
	////// return returnValue1, returnValue2,...
	// }

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
