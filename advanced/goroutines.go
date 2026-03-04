package advanced

import (
	"fmt"
	"time"
)

// ********************************************************************************************

// Goroutines : son solo funciones que viven en el metodo
// main (main thread) y se ejecutan tras bambalinas y se unen
// de vuelta con el main thread cuando una funcion
// finaliza/Lista para retornar cualquier valor

// Goroutines: No detienen el flujo del programa y no son bloqueantes (blocking)

// ********************************************************************************************

func main() {
	var err error

	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After say Hello function.")
	time.Sleep(2 * time.Second)

	go func() {
		err = doWork()
	}()

	// err = go doWork() // De esta forma no es aceptado(this is not accepted)
	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("work completed successfully")
	}
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number: ", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcdefg" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)

	return fmt.Errorf("an Error occured in doWork.")
}
