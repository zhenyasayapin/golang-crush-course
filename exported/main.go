package main

import (
	"exported/staff"
	"log"
)

var employees = []staff.Employee{
	{FirstName: "Alice", LastName: "Johnson", FullTime: true, Salary: 70000},
	{FirstName: "Bob", LastName: "Smith", FullTime: false, Salary: 40000},
	{FirstName: "Charlie", LastName: "Brown", FullTime: true, Salary: 80000},
	{FirstName: "Diana", LastName: "Prince", FullTime: true, Salary: 90000},
	{FirstName: "Ethan", LastName: "Hunt", FullTime: false, Salary: 50000},
}

func main() {
	office := staff.Office{
		AllStaff: employees,
	}

	log.Print(office.All())
	log.Println("Overpaid:", office.Overpaid())
	log.Println("Underpaid:", office.Underpaid())
}
