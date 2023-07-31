package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "1",
						Name: "Jhon Doe",
						Age:  25,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					DNI:  "1",
					Name: "Jhon Doe",
					Age:  35,
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, error := GetFullTimeEmployee(test.id, test.dni)
		if error != nil {
			t.Errorf("Error When getting Employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
