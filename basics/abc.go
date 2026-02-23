package basics

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World Hola desde Go!!!")
	})
	http.ListenAndServe(":8888", nil)
}
