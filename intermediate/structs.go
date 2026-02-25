package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func main() {

	p1 := Person{
		firstName: "Simon",
		lastName:  "Golden",
		age:       13,
		address: Address{
			city:    "Santiago",
			country: "Chile",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "2953268",
			cell: "+5697452819",
		},
	}

	p2 := Person{
		firstName: "Vigo",
		age:       8,
	}

	p2.address.city = "Puente Alto City"
	p2.address.country = "Chile el mejor pais de chile"
	p2.PhoneHomeCell.home = "2953268"
	p2.PhoneHomeCell.cell = "+56978797834"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.fullname())
	fmt.Println(p1.address)
	fmt.Println(p1.address.country)
	fmt.Println(p1.PhoneHomeCell.home)
	fmt.Println(p1.PhoneHomeCell.cell)
	fmt.Println(p2.address.country)
	fmt.Println(p2.PhoneHomeCell.home)
	fmt.Println(p2.PhoneHomeCell.cell)

	// Anonumous Structs

	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "correo@correo.cl",
	}
	fmt.Println(user.username)

	fmt.Println("Antes del incremento de edad: ", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("Despues del incremento de edad: ", p1.age)

}

func (p Person) fullname() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}
