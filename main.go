package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	sri := person{firstName: "Sri", lastName: "K"}
	fmt.Println(sri)
}
