package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~lo, Base64 Encoding")

	// Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decode fro base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error in decoding", err)
	}
	fmt.Println("Decoded string: ", string(decoded))

	// URL safe, avoid '/' and '+'

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe encoded:", urlSafeEncoded)

	urlSafeDecoded, err := base64.URLEncoding.DecodeString(urlSafeEncoded)
	if err != nil {
		fmt.Println("Error in decoding", err)
	}
	fmt.Println("URL Safe encoded:", string(urlSafeDecoded))

}
