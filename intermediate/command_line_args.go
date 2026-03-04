package intermediate

import (
	"flag"
	"fmt"
)

func main() {

	// fmt.Println("Command:", os.Args[0])

	// for i, arg := range os.Args {
	// 	fmt.Printf("Argument %d: %v\n", i, arg)
	// }

	// define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "Simon", "Nombre del usuario")
	flag.IntVar(&age, "age", 18, "Edad del usuario")
	flag.BoolVar(&male, "male", true, "Genero del usuario")

	flag.Parse()
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)
}
