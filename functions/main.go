package main

import (
	"fmt"
	"time"
)

// func double
func double(number int) int {
	return number * 2
}

func main() {
	/*x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)
	z := 10
	fmt.Println(double(z))*/
	c := make(chan int)
	go func() {
		fmt.Println("starting function anonimus")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1

	}()
	<-c
}
