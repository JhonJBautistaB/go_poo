package main

import (
	"fmt"
	"time"
)

func DoSomenthing(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	durations1 := 4 * time.Second
	durations2 := 2 * time.Second

	go DoSomenthing(durations1, canal1, 1)
	go DoSomenthing(durations2, canal2, 2)

	//fmt.Println(<-canal1)
	//fmt.Println(<-canal2)

	for i := 0; i < 2; i++ {
		select {
		case canalMsq1 := <-canal1:
			fmt.Println(canalMsq1)
		case canalMsq2 := <-canal2:
			fmt.Println(canalMsq2)
		}
	}
}
