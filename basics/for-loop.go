package basics

import "fmt"

func main() {

	//Simple iteration over range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Iterate over collection
	// numbers := []int{1, 2, 3, 4, 5, 6}

	// for index, value := range numbers {
	// 	fmt.Printf("Index: %v, Value: %v\n", index, value)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue // continue the loop but skip the rest of Lines/statments
	// 	}
	// 	fmt.Println("Odd Number:", i)
	// 	if i == 5 {
	// 		break // break out of the Loop
	// 	}

	// }

	// rows := 5

	// // Outer Loop
	// for i := 1; i <= rows; i++ {
	// 	// Inner Loop for spaces before stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	// inner loop for stars
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // Move to the next Line
	// }

	for i := range 10 {
		i++
		fmt.Println(i)
	}
	fmt.Println("We have a lift off!")
}
