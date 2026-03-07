package advanced

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex // CONSISTENCIA DE LOS RESULTADOS, sin mutex pierde la consistencia

	numGoroutines := 5
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()
	fmt.Printf("Final counter value %d\n", counter)
	fmt.Println(runtime.NumCPU())
}

// type counter struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (c *counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// }

// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }

// func main() {

// 	var wg sync.WaitGroup
// 	counter := &counter{}
// 	numGoroutines := 10

// 	// wg.Add(numGoroutines)
// 	for range numGoroutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 1000 {
// 				counter.increment()
// 				// counter.count++ // con el incremento manual el resultado no es consistente
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Printf("Final counter value: %d\n", counter.getValue())

// }
