package basics

import "fmt"

func main() {

	// Switch Statment in go is: Switch -> case -> default (fallthrough)
	// switch expression {
	// case value1:
	// Codigo para ser ejecutado si la exprecion es igual a value1
	// fallthrough
	// case value2:
	// Codigo para ser ejecutado si la exprecion es igual a value2
	// case value3:
	// Codigo para ser ejecutado si la exprecion es igual a value3
	// default:
	// Codigo para ser ejecutado si la expresion no hace match con ningun valor
	// }

	// Switch Statment in other languages is: Switch -> case -> break -> default
	// switch expression {
	// case value1:
	// Codigo para ser ejecutado si la exprecion es igual a value1
	// break
	// case value2:
	// Codigo para ser ejecutado si la exprecion es igual a value2
	// break
	// case value3:
	// Codigo para ser ejecutado si la exprecion es igual a value3
	// break
	// default:
	// Codigo para ser ejecutado si la expresion no hace match con ningun valor
	// }

	// fruit := "pineapple"
	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple")
	// case "banana":
	// 	fmt.Println("It's an banana")
	// default:
	// 	fmt.Println("Uknown fruit!")
	// }

	// Multiple conditions

	// day := "Saturday"
	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thusrday", "Friday":
	// 	fmt.Println("Is's a weekday")
	// case "Saturday", "Sunday":
	// 	fmt.Println("It's a weekend.")
	// default:
	// 	fmt.Println("Invalid day.")

	// }

	// number := 15
	//
	// switch {
	// case number < 10:
	// fmt.Println("NUmber es menor que 10")
	// case number >= 10 && number < 20:
	// fmt.Println("Numero esta entre 10 y 19")
	// default:
	// fmt.Println("Numero es 20 o mayor")
	// }

	// num := 2

	// switch {
	// case num > 1:
	// 	fmt.Println("Numero es mayor que 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Numero es 2")
	// default:
	// fmt.Println("No es 2")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hola gil desde GO")
	checkType(true)
}

func checkType(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("Es un integer")
	case int32:
		fmt.Println("Es un integer 32")
	case float64:
		fmt.Println("Es un float")
	case string:
		fmt.Println("Es una cadena de caracteres")
	default:
		fmt.Println("Uknown Type")

	}
}
