package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	//alex := person{firstname: "Alex", lastname: "Anderson"}
	//fmt.Println(alex)

	// adding a different way of assinging value to struct

	var alex person
	alex.firstname = "Alex"
	alex.lastname = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)

}
