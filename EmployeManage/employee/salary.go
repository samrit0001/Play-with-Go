package employee

func CalculateAvgSalary(employees ...Employee) float64 {
	total := 0.0
	for _, emp := range employees {
		total += emp.Salary
	}

	return total / float64(len(employees))
}
