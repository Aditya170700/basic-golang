package main

import "fmt"

func main() {
	name, age := getPerson()

	fmt.Println(name, age)
}

func getPerson() (name string, age int) {
	name = "Aditya"
	age = 23

	return name, age
}
