package main

import "fmt"

// funciones variadicas
func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// retorno con nombres
/*func getValues(x int) (int, int, int) {
	return x * 2, x * 3, x * 4
}*/

func getValues(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	printNames("pepo", "pipo", "papo")
	fmt.Println(getValues(2))
}
