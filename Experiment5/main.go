package main

import "fmt"

// modifyValue accepts a pointer and modifies the original variable.
func modifyValue(num *int) {
	*num = 100
}

type Student struct {
	Name string
	Age  int
}

func main() {

	// --------------------------------
	// PART 1: Referencing & Dereferencing
	// --------------------------------

	number := 50

	fmt.Println("Original Value:", number)
	fmt.Println("Address of number:", &number)
	fmt.Println("Value using pointer:", *(&number))

	// --------------------------------
	// PART 2: Modify variable using pointer
	// --------------------------------

	fmt.Println("\nBefore Function Call:", number)

	modifyValue(&number)

	fmt.Println("After Function Call:", number)

	// --------------------------------
	// PART 3: Struct using new()
	// --------------------------------

	student := new(Student)

	student.Name = "Rahul"
	student.Age = 21

	fmt.Println("\nStudent Details:")
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)

	// Modify struct fields through pointer
	student.Name = "Aman"
	student.Age = 22

	fmt.Println("\nAfter Modification:")
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
}
