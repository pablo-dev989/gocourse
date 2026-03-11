package advanced

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main() {

	var pool = sync.Pool{
		// New: func() interface{} {
		// 	fmt.Println("Creating a new Person.")
		// 	return &person{}
		// },
	}
	pool.Put(&person{name: "john", age: 81})

	// Get an pbject from the pool
	person1 := pool.Get().(*person)
	// person1.name = "Simon"
	// person1.age = 13
	// fmt.Println("Got Person1:", person1)

	fmt.Printf("Person1 - Name: %s, Age: %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned Person2 to pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got Person again:", person2)

	person3 := pool.Get().(*person)
	fmt.Println("Got another person3:", person3)
	person3.name = "Jane"

	//Returning object to the pool again
	pool.Put(person2)
	pool.Put(person3)
	fmt.Println("Returned Person to pool ")

	person4 := pool.Get().(*person)
	fmt.Println("Got person4:", person4)

	person5 := pool.Get().(*person)
	fmt.Println("Got person5:", person5)
}
