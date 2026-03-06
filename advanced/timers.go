package advanced

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("Timer1 expired")
	case <-timer2.C:
		fmt.Println("Timer2 expired")
	}
}

// ======= SCHEDULING DELAYED OPERATIONS
// func main() {
// 	timer := time.NewTimer(2 * time.Second) // Non blocking timer starts

// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed operation exceuted")
// 	}()

// 	fmt.Println("Waiting...")
// 	time.Sleep(3 * time.Second) // blocking time starts
// 	fmt.Println("End of Program")
// }

//===========TIme Out
// func longRunningOperation() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeout := time.After(3 * time.Second)
// 	done := make(chan bool)
// 	go func() {
// 		longRunningOperation()
// 		done <- true
// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("Operation times out")
// 	case <-done:
// 		fmt.Println("Operation completed")
// 	}
// }

// =========== Basic timer use
// func main() {
// 	// time.Sleep(time.Second)
// 	fmt.Println("Starting App.")
// 	timer := time.NewTimer(2 * time.Second)
// 	fmt.Println("Waiting for timer.c")
// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Time stopped")
// 	}
// 	timer.Reset(time.Second)
// 	fmt.Println("Timer reset")
// 	<-timer.C // blocking in nature
// 	fmt.Println("Timer expired")
// }
