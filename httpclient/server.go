package httpclient

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hola Simon!!!")
	})

	// const serverAddr string = "127.0.0.1:3000"
	const port string = ":8080"

	fmt.Println("Server Listening on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("error starting server", err)
	}

}
