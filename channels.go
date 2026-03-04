package main

import (
	"fmt"
	"time"
)

func channels() {

	// variable := make(chan type) '<-' operator
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString // blocking porque esta continuamente intentando recibir valores, esta listo para recibir un flujo continuo de datos.
		greeting <- "World"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()

	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)
	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of Program")

}
