package intermediate

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}

	defer file.Close()

	// write data to file

	data := []byte("Hola Simon, te saludo desde un archivo escrito desde Golang!!!\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writting to file:", err)
		return
	}
	fmt.Println("Data has been written to file successsfuly!")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error creating file.", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hola Go!\n")
	if err != nil {
		fmt.Println("Error writing to file", err)
		return
	}

}
