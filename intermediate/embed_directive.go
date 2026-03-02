package intermediate

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed example.txt
var content string

//go:embed basics
var basicsFolder embed.FS

func checkError(err error) {
	if err != nil {
		panic(err)
		// fmt.println()
	}
}

func main() {

	fmt.Println("Embedded content:", content)
	content, err := basicsFolder.ReadFile("basics/hello.txt")
	checkError(err)
	fmt.Println("Embedded file content:", string(content))

	err = fs.WalkDir(basicsFolder, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(path)
		return nil

	})
	checkError(err)
}
