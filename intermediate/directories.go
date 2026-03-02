package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		// fmt.println()
	}
}

func main() {

	// err := os.Mkdir("subdir", 0755)
	// checkError(err)
	// Desde ahora usaremos solo una linea de codigo para validar si el error exite
	// checkError(os.Mkdir("subdir", 0755))

	// // defer os.RemoveAll("subdir")
	// os.WriteFile("subdir/file", []byte(""), 0755)

	// checkError(os.MkdirAll("subdir1/parent/child", 0755))
	// checkError(os.MkdirAll("subdir1/parent/child1", 0755))
	// checkError(os.MkdirAll("subdir1/parent/child2", 0755))
	// checkError(os.MkdirAll("subdir1/parent/child3", 0755))
	// checkError(os.MkdirAll("subdir1/parent/child4", 0755))
	// os.WriteFile("subdir1/parent/file", []byte(""), 0755)
	// os.WriteFile("subdir1/parent/child/file", []byte(""), 0755)

	result, err := os.ReadDir("subdir1/parent")
	checkError(err)
	// fmt.Println(result)
	fmt.Println("Reading subdir/parent")
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	checkError(os.Chdir("subdir1/parent/child"))

	result, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading subdir/parent/child")
	for _, entry := range result {
		fmt.Println(entry)
	}

	checkError(os.Chdir("../../.."))
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)

	// filepath.Walk and filepart.WalkDir

	pathfile := "subdir1/parent/child"
	// pathfile := "."

	fmt.Println("Walking Directory")
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		fmt.Println(path)
		return nil
	})
	checkError(err)

	// checkError(os.RemoveAll("subdir1"))
	checkError(os.Remove("subdir1"))

}
