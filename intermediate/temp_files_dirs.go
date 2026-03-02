package intermediate

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// tempFileName := "temporaryFile"

	// tempFile, err := os.CreateTemp("", tempFileName)
	// checkError(err)

	// fmt.Println("Temporary file creted:", tempFile.Name())

	// defer os.Remove(tempFile.Name())
	// defer tempFile.Close()

	tempDir, err := os.MkdirTemp("", "GoCourseTemDir")
	checkError(err)

	defer os.RemoveAll(tempDir)

	fmt.Println("Temporary directory created:", tempDir)

}
