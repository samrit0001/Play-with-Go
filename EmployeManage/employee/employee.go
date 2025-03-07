package employee

import (
	"fmt"
)

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

// e is receiver, so Dispaly is method of Employee
// so Display can use the variables of Employee using the receiver variable e
func (e Employee) Display() {
	fmt.Printf("Name : %s , Age : %d , Salary : %.2f ", e.Name, e.Age, e.Salary)
}
