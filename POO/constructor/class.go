package main

import "fmt"

type Employee struct {
	id        int
	name      string
	vacations bool
}

// funcion que emula el metodo constructor
func NewEmployee(id int, name string, vacations bool) *Employee {
	return &Employee{
		id:        id,
		name:      name,
		vacations: vacations,
	}
}

func main() {
	// forma 1 de instanciar
	e := Employee{}
	fmt.Println(e)

	// forma 2 de instanciar
	e2 := Employee{
		id:        1,
		name:      "Name",
		vacations: true,
	}
	fmt.Println(e2)

	// forma 3 de instanciar
	e3 := new(Employee)
	fmt.Println(e3)
	e3.id = 2
	e3.name = "Name 2"
	e3.vacations = false
	fmt.Println(*e3)

	//forma 4 a traves de una función constructora
	e4 := NewEmployee(4, "name4", true)
	fmt.Printf("%v\n", *e4)

}
