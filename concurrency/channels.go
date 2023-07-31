package main

import "fmt"

func main() {
	c := make(chan int, 4)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
