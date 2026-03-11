package advanced

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	// "Esto no se repetirá sin importar cuántas veces llames a esta función usando once.Do."
	fmt.Println("This will not be repeated no matter how many times you call this function using once.Do")
}

func main() {

	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine # ", i)
			once.Do(initialize)
		}()
	}
	wg.Wait()

}
