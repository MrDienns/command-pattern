package main

import "fmt"

func main() {

	person := &Person{name: "Dennis"}
	fmt.Printf("Person name: %v\n", person.Name())

	command := CommandSetName{
		person: person,
		name:   "Martin",
	}
	command.Execute()

	fmt.Printf("Person name: %v\n", person.Name())
}
