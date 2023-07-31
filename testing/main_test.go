/* COMANDOS MAS USADOS PARA TESTEAR
* go test -> realiza test y muestra OK y el tiemppo del testeo
* go test cover -> realiza el test mostrando el coverage en % y tiempo de ejecución
* go test -coverprofile=coverage.out -> crea un archivo del test para ejecutar metricas de coverage
* go tool cover -func=coverage.out -> a partir del archivo de coverage devuelve el % de coverage por cada funcion testeada
* go tool cover -html=coverage.out -> a partir del archivo de coverage devuelve un archivo html donde se identifica hasta donde esta testeada una func
* go test -cpuprofile=cpu.out -> crea un archivo del test y guarda el uso de cpu de un test a una función
* go tool pprof cpu.out -> herramienta que permite a través de unos comandos ver metricas de comportamiento de testing
* top -> muestra todas las funciones testeadas el porcentaje de uso y consumo
* list <funcion> -> muestra los tiempos donde mas se esta demorando una función
 */

package main

import "testing"

/*
	func TestSum(t *testing.T) {
		total := Sum(5, 3)
		if total != 10 {
			t.Errorf("Sum was incorrect, got %d expect %d", total, 10)
		}
	}
*/
func TestSum(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expect %d", total, item.n)
		}
	}

}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{2, 5, 5},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d expect %d", max, item.n)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorret, got %d expected %d", fib, item.n)
		}
	}
}
