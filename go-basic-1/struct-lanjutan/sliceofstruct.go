package main

import "fmt"

type Employee struct {
	Name        string
	Departement string
}

type Employees []Employee

func main() {

	var employees Employees
	fmt.Println(employees)

	// employees[0].Name = "Employee A" ini akan error dikarenakan saat deklarasi awal, employees tidak mempunyai element sama sekali (slice kosong)

	var emp1 = Employee{Name: "Emp1", Departement: "Tech"}
	employees = append(employees, emp1)

	fmt.Println(employees)

	employees = append(employees, Employee{Name: "Emp2", Departement: "Finance"})
	fmt.Println(employees)

}
