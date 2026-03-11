package advanced

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorted struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (s *personSorted) Len() int {
	return len(s.people)
}

func (s *personSorted) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorted) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort(people []Person) {
	ps := &personSorted{
		people: people,
		by:     by,
	}
	sort.Sort(ps)

}

func main() {

	people := []Person{
		{"Pablo", 45},
		{"Simon", 13},
		{"Antonio", 67},
	}
	fmt.Println("Unsorted (original):", people)
	// sort.Sort(ByAge(people))
	ageAsc := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}
	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	lenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}
	By(ageAsc).Sort(people)
	fmt.Println("Sorted by Age (ascending):", people)
	By(ageDesc).Sort(people)
	fmt.Println("Sorted by Age (descending):", people)
	By(name).Sort(people)
	fmt.Println("Sorted by Name:", people)
	By(lenName).Sort(people)
	fmt.Println("Sorted by Length of Name:", people)

	// ==== SORT SLICE
	stringSlice := []string{"banana", "apple", "cherry", "grapes", "guava"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})

	fmt.Println("Sorted by last character:", stringSlice)

	// numbers := []int{5, 3, 4, 1, 2}
	// sort.Ints(numbers)
	// fmt.Println("Sorted Numbers: ", numbers)

	// stringSlice := []string{"Pablo", "Simon", "Antonio", "Victor", "Felipe"}
	// sort.Strings(stringSlice)
	// fmt.Println("Sorted String: ", stringSlice)

	// type ByAge []Person
	// type ByName []Person

	// func (a ByAge) Len() int {
	// 	return len(a)
	// }

	// func (a ByAge) Less(i, j int) bool {
	// 	return a[i].Age < a[j].Age
	// }

	// func (a ByAge) Swap(i, j int) {
	// 	a[i], a[j] = a[j], a[i]
	// }

	// func (a ByName) Len() int {
	// 	return len(a)
	// }

	// func (a ByName) Less(i, j int) bool {
	// 	return a[i].Name < a[j].Name
	// }

	// func (a ByName) Swap(i, j int) {
	// 	a[i], a[j] = a[j], a[i]
	// }

}
