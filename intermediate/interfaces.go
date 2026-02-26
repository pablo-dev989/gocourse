package intermediate

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, heigth float64
}
type rect1 struct {
	width, heigth float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.heigth * r.width
}

func (r rect1) area() float64 {
	return r.heigth * r.width
}

func (r rect1) perim() float64 {
	return 2 * (r.heigth + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.heigth + r.width)
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	r := rect{width: 3, heigth: 4}
	c := circle{radius: 5}
	// r1 := rect1{width: 3, heigth: 4}

	measure(r)
	measure(c)
	// measure(r1)

	myPrinter(1, "Simon", 16.8, true)

	printType(9)
	printType("Es es un string")
	printType(false)

}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Is an Integer")
	case string:
		fmt.Println("Is an String")
	default:
		fmt.Println("Type Unknown")

	}
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
