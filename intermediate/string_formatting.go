package intermediate

import "fmt"

func main() {

	num := 42
	fmt.Printf("%05d\n", num)

	// message := "Hello aca ponemos un string mas largo ..."
	message := "Hello"
	fmt.Printf("|%10s|\n", message)  // --->> Justifica a la derecha
	fmt.Printf("|%-10s|\n", message) // <<--- Justifica a la Izquierda

	message1 := "Hello \nWorld!"
	message2 := `Hello \nWorld`

	fmt.Println(message1)
	fmt.Println(message2)

	sqlquery := `
	SELECT *
	  FROM users
	 WHERE age > 30
	`

	fmt.Println(sqlquery)
}
