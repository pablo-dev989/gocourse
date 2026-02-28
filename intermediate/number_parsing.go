package intermediate

import (
	"fmt"
	"strconv"
)

func main() {

	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value.", err)
	}
	fmt.Println("Parsed Integer:", num)
	fmt.Println("Parsed Integer:", num+1)

	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing value.", err)
	}
	fmt.Println("Parsed Integer:", numistr)
	fmt.Println("Parsed Integer:", numistr+1)

	floatStr := "3.14"
	floatval, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error parsing value.", err)
	}
	fmt.Printf("Float Value is: %.2f\n", floatval)

	binaryStr := "1010"
	decimal, err := strconv.ParseInt(binaryStr, 2, 8)
	if err != nil {
		fmt.Println("Error parsing value ACA.", err)
	}
	fmt.Println("Parsed binary to decimal value ACA:", decimal)

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing value.", err)
	}
	fmt.Println("Parsed binary to hex value:", hex)

	invalidNum := "456abc"
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing value.", err)
	}
	fmt.Println("Parsed invalid number:", invalidParse)
}
