package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
	return
}

func (e *Employee) SetName(name string) {
	e.name = name
	return
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Println(e)
	e.id = 1
	e.name = "my name"
	fmt.Println(e)

	e.SetId(5)
	fmt.Println(e)

	e.SetName("llego aquí")
	fmt.Println(e)

	fmt.Println("getId->", e.GetId())
	fmt.Println("getName->", e.GetName())
}
