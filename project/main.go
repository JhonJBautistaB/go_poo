package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Class Job
type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

// Class Worker
type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

// Method class Worker
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id %d Started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d Finished with result %d\n", w.Id, fib)
			case <-w.QuitChan:
				fmt.Printf("Worker with id %d Stopped\n", w.Id)
			}
		}
	}()
}

// Method class Worker
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// funcion para crear un nuevo worker
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

// Class dispatcher
type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

// Method class dispatcher
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

// Method class dispatcher
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

// funcion para crear un nuevo dispatcher
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

// functions fibonacci for app used
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Handle the request from the server
func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Add("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the request
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	// Create the job
	job := Job{
		Name:   name,
		Delay:  delay,
		Number: value,
	}

	// Add the job to the queue
	jobQueue <- job // send the job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()

	// http://localhost:8081/fib
	http.HandleFunc(
		"/fib",
		func(w http.ResponseWriter, r *http.Request) {
			RequestHandler(w, r, jobQueue)
		})
	log.Fatal(http.ListenAndServe(port, nil))
}
