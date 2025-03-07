package main

import (
	"app/employee"
	"fmt"
)

func main() {

	var employees []employee.Employee

	newEmployee := employee.Employee{Name: "amrit", Age: 25, Salary: 5000.0}

	employees = append(employees, newEmployee)

	anotherEmployee := employee.Employee{Name: "Motti", Age: 23, Salary: 10000.0}
	employees = append(employees, anotherEmployee)

	fmt.Println("Employee Details")

	for _, emp := range employees {
		emp.Display()
	}

	averageSalary := employee.CalculateAvgSalary(employees...)
	fmt.Printf("average salary is %.2f\n", averageSalary)
}
