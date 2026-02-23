package basics

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("Deferred Statement")

	fmt.Println("Start the main funbction")

	// Exit withn status code of 1
	os.Exit(1)

	//This will never be executed
	fmt.Println("End of main function")

}
