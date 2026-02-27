package intermediate

import (
	"fmt"
	"strings"
)

func main() {

	// str := "Hello Go!"

	// fmt.Println(len(str))

	// str1 := "Hello"
	// str2 := "World!"
	// result := str1 + " " + str2
	// fmt.Println(result)

	// fmt.Println(str[0])
	// fmt.Println(str[1:7])

	// // String Conversion
	// num := 18
	// str3 := strconv.Itoa(num)
	// fmt.Println(len(str3))

	// // string splitting
	// fruits := "apple, orange, banana"
	// fruits1 := "apple-orange-banana"

	// part := strings.Split(fruits, ", ")
	// part1 := strings.Split(fruits1, "-")

	// fmt.Println(part)
	// fmt.Println(part1)

	// countries := []string{"Germany", "France", "Italy", "Chile", "Mexico"}
	// joined := strings.Join(countries, ", ")
	// fmt.Println(joined)

	// fmt.Println(strings.Contains(str, "Go"))

	// replace := strings.Replace(str, "Go", "World", 1)
	// fmt.Println(replace)

	// strwspace := " Hola Mundo! "
	// fmt.Println(strwspace)
	// fmt.Println(strings.TrimSpace(strwspace))

	// fmt.Println(strings.ToLower(strwspace))
	// fmt.Println(strings.ToUpper(strwspace))

	// fmt.Println(strings.Repeat("foo", 3))

	// fmt.Println(strings.Count("Hello", "l"))
	// fmt.Println(strings.HasPrefix("Hello", "he"))
	// fmt.Println(strings.HasSuffix("Hello", "lo"))
	// fmt.Println(strings.HasSuffix("Hello", "la"))

	// str5 := "Hel1lo, 123 Go 11!"
	// re := regexp.MustCompile(`\d+`)
	// matches := re.FindAllString(str5, -1)
	// fmt.Println(matches)

	// str6 := "Hello, 世界"
	// fmt.Println(utf8.RuneCountInString(str6))

	// STRING BUILDER
	var builder strings.Builder

	// write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	//Conver builder to a string
	result := builder.String()
	fmt.Println(result)

	// Using writerune to add character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	//Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)

}
