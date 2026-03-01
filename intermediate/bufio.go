package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reader := bufio.NewReader(strings.NewReader("Hello, bufio packageee!\nHow are you doing?"))

	// // Reading byte slice
	// data := make([]byte, 20)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("Error encoding:", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading string", err)
	// 	return
	// }
	// fmt.Println("Read string: ", line)

	writer := bufio.NewWriter(os.Stdout)

	// writing byte slice

	data := []byte("Hello, bufio Package!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error Flushing writter:", err)
		return
	}

	// writting string
	str := "This is a string.\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n)

	// flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error Flushing writter:", err)
		return
	}

}
