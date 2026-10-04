package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

// ReadData reads Person details from user input.
func (p *Person) ReadData() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter Age: ")
	ageInput, _ := reader.ReadString('\n')
	age, _ := strconv.Atoi(strings.TrimSpace(ageInput))
	p.Age = age

	fmt.Print("Enter Job: ")
	job, _ := reader.ReadString('\n')
	p.Job = strings.TrimSpace(job)

	fmt.Print("Enter Salary: ")
	salaryInput, _ := reader.ReadString('\n')
	salary, _ := strconv.ParseFloat(strings.TrimSpace(salaryInput), 64)
	p.Salary = salary
}

// Display prints all Person details.
func (p Person) Display() {
	fmt.Println("\n--- Person Details ---")
	fmt.Println("Name   :", p.Name)
	fmt.Println("Age    :", p.Age)
	fmt.Println("Job    :", p.Job)
	fmt.Printf("Salary : %.2f\n", p.Salary)
}

func main() {
	var person1 Person
	var person2 Person

	fmt.Println("Enter details for Person 1:")
	person1.ReadData()

	fmt.Println("\nEnter details for Person 2:")
	person2.ReadData()

	person1.Display()
	person2.Display()
}
