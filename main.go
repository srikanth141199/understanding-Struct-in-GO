package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contractInfo
}

type contractInfo struct {
	email   string
	zipCode int
}

func main() {
	// sri := person{firstName: "Sri", lastName: "K"}
	// fmt.Println(sri)
	// //diff way to define struct
	// var ram person
	// ram.firstName = "Ram"
	// ram.lastName = "R"
	// fmt.Println(ram)
	// fmt.Printf("%+v", ram)
	// //diff way
	// Krish := person{"Krish", "K", contractInfo{"sri@sri.com", 1234}}
	// fmt.Println(Krish)

	Krish := person{
		"Krish",
		"K",
		contractInfo{
			"sri@sri",
			1234,
		},
	}

	KrishPointer := &Krish
	KrishPointer.updateName("Krishna")
	Krish.print()
}

func (personAsPointer *person) updateName(newFirstName string) {
	(*personAsPointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
