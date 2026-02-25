package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r Rectangle)

func main() {

	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("El area del rectangulo con un ancho de 9 y un largo de 10 es:", area)

}
