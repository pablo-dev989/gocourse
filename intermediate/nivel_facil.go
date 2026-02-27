package intermediate

import "fmt"

func main() {

	autos := []string{"BYD", "Fiat", "Ferrari"}

	moderno := []string{}

	for _, v := range autos {
		if v == "Ferrari" || v == "Fiat" {
			moderno = append(moderno, v)
		}
	}
	fmt.Println(moderno)

}
