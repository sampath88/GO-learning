package main

import "fmt"

//package level variable
const isActive = true

func main() {
	// var name string 	//default empty string
	// var age int // default value = 0
	// var isActive bool  //default false
	// var salary float32	// default	0
	// var u uint			// default 	0

	// fmt.Println("name: ",name)
	// fmt.Println("age: ",age)
	// fmt.Println("isActive: ",isActive)
	// fmt.Println("salary: ",salary)
	// fmt.Println("u: ",u)

	var x int8 = 127
	fmt.Println(x)

	var name string = "This is sampath"
	fmt.Println(name)

	//short hand declaration of variables
	age := 18
	myName := "Sampath"

	fmt.Printf("My name is %v. I am %v years old", myName, age)

	//constants
	const country = "India"
	fmt.Printf("\nI am from %v\n", country)

	fmt.Printf("I am active %v\n",isActive)
}
