package main

import (
	"fmt"
	"time"
)

type Employee struct {
	id      int
	name    string
	country string
	created time.Time
}

func main() {
	employee()
}

func employee() {
	var employeeOne Employee
	employeeTwo := new(Employee)

	employeeOne.id = 1
	employeeOne.country = "UK"
	employeeOne.name = "John Smith"
	employeeOne.created = time.Now()

	employeeTwo.id = 2
	employeeTwo.name = "William Jefferson"
	employeeTwo.country = "United States"
	employeeTwo.created = time.Now()

	fmt.Println("<===== First Employee =====>")
	fmt.Println("Id: ", employeeOne.id)
	fmt.Println("Name: ", employeeOne.name)
	fmt.Println("Country: ", employeeOne.country)
	fmt.Println("Time Created: ", employeeOne.created)

	fmt.Println()
	fmt.Println("<===== Second Employee =====>")
	fmt.Println("Id: ", employeeTwo.id)
	fmt.Println("Name: ", employeeTwo.name)
	fmt.Println("Country: ", employeeTwo.country)
	fmt.Println("Time Created: ", employeeTwo.created)
}
