package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   float64
	FullTime bool
}

func main() {
	jack := Employee{
		Name:     "Jack",
		Age:      35,
		Salary:   60000.0,
		FullTime: true,
	}

	jill := Employee{
		Name:     "Jill",
		Age:      25,
		Salary:   65000.0,
		FullTime: false,
	}

	for _, employee := range []Employee{jack, jill} {
		if employee.Age > 30 {
			fmt.Println(employee.Name, "is older than 30")
		}

		if employee.Age <= 30 {
			fmt.Println(employee.Name, "is 30 or younger")
		}

		if employee.Salary >= 60000.0 && employee.FullTime {
			fmt.Println(employee.Name, "earns at least 60000 and is a full-time employee")
		}

		if employee.Age >= 30 || employee.Salary > 30000.0 && employee.FullTime {
			fmt.Println(employee.Name, "is either at least 30 years old, or earns more than 30000 and is a full-time employee")
		}
	}
}
