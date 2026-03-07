package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate proccessing of ticket requets
func ticketProccessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Proccessing %d ticket(s) of personID %d with total cost %d\n", req.numTickets, req.personID, req.cost)
		// simulate proccessing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	// start ticket proccessor/ticket
	for range 3 {
		go ticketProccessor(ticketRequests, ticketResults)
	}

	// send ticket reqsuests
	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for person ID %d proccessed successfully!\n", <-ticketResults)
	}

}

//==== BASIC WORKER POOL PATTERN
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d proccessing task %d\n", id, task)
// 		// Simulate work
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }

// func main() {

// 	numWorkers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	// Create Workers
// 	for i := range numWorkers {
// 		go worker(i, tasks, results)
// 	}

// 	// Send values to the tasks channel
// 	for i := range numJobs {
// 		tasks <- i
// 	}

// 	// Collects the result
// 	for range numJobs {
// 		result := <-results
// 		fmt.Println("Result:", result)
// 	}

// }
