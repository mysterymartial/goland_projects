package main

import "fmt"

type Employee struct {
	name string
	id   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s is an Employee", e.name, e.id)

}

type Manager struct {
	Employee
	Report []Employee
}
