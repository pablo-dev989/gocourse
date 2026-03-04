package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name" db:"firstn" xml:"first"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age,omitempty"`
}

func main() {

	// person := Person{
	// 	FirstName: "Simon",
	// 	LastName:  "",
	// 	Age:       13,
	// }
	person := Person{
		FirstName: "Simon",
		Age:       13,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalln("Error Marshalling struct:", err)
	}
	fmt.Println(string(jsonData))

}
