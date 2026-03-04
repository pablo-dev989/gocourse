package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	Email   string   `xml:"-"`
	Address Address  `xml:"address"`
	// Email   string   `xml:"email"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {

	// person := Person{Name: "Simon", Age: 30, City: "Santiago", Email: "simon.golden@gmail.com"}
	person := Person{Name: "Simon", Email: "simon.golden@gmail.com", Address: Address{City: "Santiago", State: "Region Metrolpolitana"}}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error Marshalling xml:", err)
	}

	fmt.Println(string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error Marshalling xml:", err)
	}

	fmt.Println(string(xmlData1))

	// xmlRaw := `<person><name>Simon</name><age>13</age></person>`
	xmlRaw := `<person><name>Simon</name><age>13</age><address><city>Santiago</city><state>Region Metropolitana</state></address></person>`

	var personxml Person
	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatalln("Error Unmarshalling XML:", err)
	}
	fmt.Println(personxml)
	fmt.Println("Local String", personxml.XMLName.Local)
	fmt.Println("Namespace", personxml.XMLName.Space)

	book := Book{
		ISBN:       "978180345623",
		Title:      "Go Bootcamp",
		Author:     "Ashish",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attribute",
	}

	xmlDataAttr, err := xml.MarshalIndent(book, "", "   ")
	if err != nil {
		log.Fatalln("Error marshaling data:", err)
	}
	fmt.Println(string(xmlDataAttr))
}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}
