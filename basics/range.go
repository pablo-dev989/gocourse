package basics

import "fmt"

func main() {

	message := "Hola mundo."

	for i, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

}
