// ===================================================================================

// Go’s structs are typed collections of fields.
// It is a user-defined type that allows to group/combine items of possibly different types into a single type.
// They’re useful for grouping data together to form records.
// A structure is used mainly when you need to define a schema made of different individual fields (properties).
// 	Like a class, we can create an object from this schema
// Omitted fields will be zero-valued.

package main

import (
	"fmt"
)

type User struct {
	Name   string
	Age    int
	Salary int
}

type Employee struct {
	Name       string
	EmpId      int
	Department string
	Salary     Salary
}

type Salary struct {
	Total     int
	Basic     int
	Insurance int
}

func setUser(name string, age int) *User {
	u := User{Name: name, Age: age}
	return &u
}

func main() {

	var u User     // create a struct variable of struct User
	fmt.Println(u) //Initialize properties with zero.

	u1 := User{"Nikhil", 23, 25000}
	fmt.Println(u1)

	u2 := User{Name: "Nikhil", Age: 20, Salary: 25000}
	fmt.Println(u2)

	u3 := new(User)
	u3.Name = "Nikhil"
	u3.Age = 25
	fmt.Println(u3)

	//  To set values to struct via functions.
	fmt.Println(setUser("Mahesh", 67))

	// Creating a struct without explicitly defining it
	//  This is useful when you dont to reuse your structure

	empployee := struct {
		Name     string
		EmpId    int
		Fulltime bool
	}{
		Name:     "Suhas",
		EmpId:    1456,
		Fulltime: true,
	}

	fmt.Println("Employee Struct", empployee)

	//  Pointer to struct

	u4 := &User{
		Name: "Nikhil",
		Age:  20,
	}

	fmt.Println("Pointer to struct", *u4) // use *  to dereference struct.
	fmt.Println("Name is", (*u4).Name)    // u can write  u4.Name instead of * , Go lang will automatically derefernce.

	//  Nested Structures.

	emp := Employee{
		Name:       "Nikhil Patne",
		EmpId:      2658,
		Department: "IDM",
		Salary: Salary{
			Total:     37600,
			Basic:     15000,
			Insurance: 25000,
		},
	}

	fmt.Printf("Employee struct %+v \n", emp)

	//   retrive value of the attribute using function ( defult way emp.name )

	fmt.Println("EMP name", emp.getEmployeeName())

	//  set the value of the attribute using function  (defult way emp.name="nikhil")
	emp.setEmployeeName("Mohit")

	fmt.Println("Updated employee name after setting value", emp.getEmployeeName())

}

func (e Employee) getEmployeeName() string {
	return e.Name
}

// function to set employee name
// NOte:  U should use pointer of Struct inorder to set some value.
func (e *Employee) setEmployeeName(name string) {
	e.Name = name
}
