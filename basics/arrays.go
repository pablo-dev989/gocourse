package basics

import "fmt"

func main() {

	// Esta es la manera en que se declara un arreglo
	// var arrayName [size]elementType

	// var numbers [5]int
	// fmt.Println(numbers)

	// numbers[4] = 20
	// fmt.Println(numbers)

	// numbers[0] = 9
	// fmt.Println(numbers)

	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	// fmt.Println("Array de Frutas: ", fruits)

	// fmt.Println("Third element", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray

	// copiedArray[0] = 300

	// fmt.Println("Original Array: ", originalArray)
	// fmt.Println("Copied Array: ", copiedArray)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Elemento en el indice, ", i, " : ", numbers[i])
	// }

	// for i, v := range numbers {
	// 	fmt.Printf("Indice: %d, Valor= %d\n", i, v)
	// }

	// // underscore is blank identifier, used to store unused values
	// for _, v := range numbers {
	// 	fmt.Printf("Valor= %d\n", v)
	// }

	// a, _ := someFunction()
	// fmt.Println(a)
	// // fmt.Println(b)

	// fmt.Println("El largo del arreglo de numeros es: ", len(numbers))

	// // Comparing Arrays
	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{10, 2, 3}

	// fmt.Println("Array1 es igual al array2: ", array1 == array2)

	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Copied Array: ", copiedArray)

}

func someFunction() (int, int) {
	return 1, 2
}
