package basics

import "fmt"

func main() {

	// if condition {
	// block of code
	//}

	// age := 25

	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// }

	// if condition {
	// block of code
	//} else {
	// block of code
	//} else {
	// block of code
	//]

	// temperature := 32
	// if temperature >= 30 {
	// 	fmt.Println("Esta caluroso afuera!")
	// } else {
	// 	fmt.Println("Esta helado afuera...")
	// }

	// score := 95
	// if score >= 80 {
	// 	fmt.Println("Grade B")

	// } else if score >= 90 {
	// 	fmt.Println("Grade A")

	// } else if score >= 70 {
	// 	fmt.Println("Grade C")

	// } else {
	// 	fmt.Println("Grade D")
	// }
	// this Line will be executed after of the condition is met

	// if condition1 {
	////// code block
	///////if condition2 {
	////////////codde block 2
	///////}
	//}

	// num := 18
	// if num%2 == 0 {
	// 	if num%3 == 0 {
	// 		fmt.Println("Numero es divisible por 2 y 3.")
	// 	} else {
	// 		fmt.Println("Numero es divisible por 2 pero no 3")
	// 	}
	// } else {
	// 	fmt.Println("Number no es divisible por 2")
	// }

	// || OR
	// && AND

	if 10%2 == 0 || 5%2 == 0 {
		fmt.Println("Either 10 or 5 are even.")
	}

	if 10%2 == 0 && 5%2 == 0 {
		fmt.Println("Both 10 and 5 are even.")
	}
}
