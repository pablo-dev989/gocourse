package main

import (
	"fmt"
	"time"
)

func unbufered() {

	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(2 * time.Second)
		fmt.Println("2 seconds Goroutine finished")
	}()
	go func() {
		// ch <- 1
		time.Sleep(3 * time.Second)
		fmt.Println("3 seconds Goroutine finished")
	}()

	receiver := <-ch
	fmt.Println(receiver)
	fmt.Println("End of Program")
}
