package basics

import "fmt"

func main() {

	// For as While
	// i := 1
	// for i <= 5 {
	// 	fmt.Println("Iteration:", i)
	// 	i++
	// }

	// Infinite Loop while with a breake

	// sum := 0
	// for {
	// 	sum += 1
	// 	fmt.Println("Sum: ", sum)
	// 	if sum >= 1000 {
	// 		break
	// 	}
	// }

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Numbers: ", num)
		num++ // increment operator increases values by 1 and -- decrement operator, decreases value by 1
	}
}
