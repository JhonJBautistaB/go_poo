package main

import (
	"fmt"
	"time"
)

type Person struct {
	DNI  string
	Name string
	Age  int
}

/*
func GetPersonByDNI(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// select * from person were dni
	return Person{}, nil
}
*/
// forma equivalente de nombrar funciones en golang
var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// select * from person were dni
	return Person{}, nil
}

type Employee struct {
	Id       int
	Position string
}

/*
func GetEmployeeById(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// select * from employee were id
	return Employee{}, nil
}
*/
// forma equivalente de nombrar funciones en golang
var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// select * from employee were id
	return Employee{}, nil
}

// composición en vez de herencia
type FullTimeEmployee struct {
	Person
	Employee
}

func GetFullTimeEmployee(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee
	e, error := GetEmployeeById(id)
	if error != nil {
		return ftEmployee, error
	}
	ftEmployee.Employee = e
	p, error := GetPersonByDNI(dni)
	if error != nil {
		return ftEmployee, error
	}
	ftEmployee.Person = p

	return ftEmployee, nil
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.Name = "Pepito Perez"
	ftEmployee.Age = 35
	ftEmployee.Id = 1
	fmt.Printf("%v\n", ftEmployee)
	//GetMessage(ftEmployee)
}
