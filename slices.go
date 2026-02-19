package main

import "fmt"

func main() {

	//var sliceName[]ElementType

	// var numbers []int
	// var numbers1 = []int{1,2,3}

	// numbers2 := []int{9,8,7}

	// slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]

	fmt.Println(slice1)

}
