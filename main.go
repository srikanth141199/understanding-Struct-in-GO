package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	sri := person{firstName: "Sri", lastName: "K"}
	fmt.Println(sri)
	//diff way to define struct
	var ram person
	ram.firstName = "Ram"
	ram.lastName = "R"
	fmt.Println(ram)
	fmt.Printf("%+v", ram)
	//diff way

	Krish := person{"Krish", "K"}
	fmt.Println(Krish)
}
