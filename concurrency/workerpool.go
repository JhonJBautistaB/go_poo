package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("En el Worker id %d inicia calculo Fibonnaci para numero %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("En el Worker id %d, para el numero %d su fibonnaci es %d\n", id, job, fib)
		result <- fib
	}
}

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40, 100}
	nWorkers := 3

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}

}
