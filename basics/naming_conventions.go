package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Pascal case
	// Examples. CaculateArea, UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	// snake_case
	// Examples: user_id, first_name, http_request

	// UPPERCASE
	// Use case is Constants

	// mixedCase
	// Example: javaScript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
