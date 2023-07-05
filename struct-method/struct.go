package main

import "fmt"

type Employee struct {
	Id     string `json:"id"`
	Name   Name   `json:"name"`
	Salary Salary `json:"salary"`
}

type Salary struct {
	Basic     int `json:"basic"`
	Bonus     int `json:"bonus"`
	Allowance int `json:"allowance"`
}

type Name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var records []Employee

func (e *Employee) totalAmount() int {

	return e.Salary.Basic + e.Salary.Bonus + e.Salary.Allowance
}

func main() {

	employee := Employee{
		Id: "1",
		Name: Name{
			Firstname: "John",
			Lastname:  "Doe",
		},
		Salary: Salary{
			Basic:     2e5,
			Bonus:     10e5,
			Allowance: 1e3,
		},
	}

	records = append(records, employee)
	fmt.Println(records)

	// Accessing the TotalAmount() function for an employee
	fmt.Println("Total Amount:", records[0].totalAmount())
}
