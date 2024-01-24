package main

import "fmt"

type Engineer struct {
	name             string
	age              uint8
	specialization   string
	employmentStatus bool
}

type Trader struct {
	name             string
	age              uint8
	desk             string
	employmentStatus bool
}

func (e Engineer) getName() string {
	return e.name
}

func (t Trader) getName() string {
	return t.name
}

func (e Engineer) isEmployed() bool {
	return e.employmentStatus
}

func (t Trader) isEmployed() bool {
	return t.employmentStatus
}

func (e Engineer) printInfo() {
	fmt.Println("Engineer Information Card")
	fmt.Printf("\tName: %v\n", e.name)
	fmt.Printf("\tAge: %v\n", e.age)
	fmt.Printf("\tSpecialization: %v\n", e.specialization)
}

func (t Trader) printInfo() {
	fmt.Println("Trader Information Card")
	fmt.Printf("\tName: %v\n", t.name)
	fmt.Printf("\tAge: %v\n", t.age)
	fmt.Printf("\tDesk: %v\n", t.desk)
}

type employee interface {
	isEmployed() bool
	getName() string
	printInfo()
}

func isEmployed(e employee) {
	if e.isEmployed() {
		e.printInfo()
	} else {
		fmt.Printf("%v is not employed!\n", e.getName())
	}
}

func main() {
	var employee1 Engineer = Engineer{name: "Adam", age: 32, specialization: "Operations", employmentStatus: false}
	var employee2 Engineer = Engineer{name: "Max", age: 42, specialization: "Software Developer", employmentStatus: true}
	var employee3 Trader = Trader{name: "Charles", age: 35, desk: "Options", employmentStatus: true}

	isEmployed(employee1)
	isEmployed(employee2)
	isEmployed(employee3)
}
