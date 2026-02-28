package intermediate

import (
	"fmt"
	"math/rand"
)

func main() {

	// val := rand.New(rand.NewSource(time.Now().Unix()))

	// fmt.Println(rand.Intn(6) + 5)
	// fmt.Println(val.Intn(101))
	// fmt.Println(rand.Intn(100) + 1) // Si es que se necesita generar numeros desde el 1 al 100
	// fmt.Println(rand.Float64()) // between 0.0 and 1.0

	for {
		// Show the menu
		fmt.Println("Welcome to the Dice game!")
		fmt.Println("1. Rol the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalide choice, please enter 1 or 2.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		// Show the result
		fmt.Printf("YOu rolled a %d and %d.\n", die1, die2)
		fmt.Println("Total:", die1+die2)

		// Ask of the user to roll aagin
		fmt.Println("Do you want to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye ")
			break
		}
	}
}
