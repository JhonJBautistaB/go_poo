package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// composición en vez de herencia
type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.name)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Pepito Perez"
	ftEmployee.age = 35
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
	//GetMessage(ftEmployee)
}
