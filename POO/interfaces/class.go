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
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "full time employeee"
}

type TemporatyEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporatyEmployee) getMessage() string {
	return "Temporary employeee"
}

/*
	func GetMessage(p Person) {
		fmt.Printf("%s with age %d\n", p.name, p.name)
	}
*/
type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Pepito Perez"
	ftEmployee.age = 35
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
	//GetMessage(ftEmployee)
	tEmployee := TemporatyEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
