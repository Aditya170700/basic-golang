package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	person1 := Person{
		Name: "Aditya",
	}

	person1.sayHello("Ricki")
}

func (p Person) sayHello(toName string) {
	fmt.Println("Hello", toName, "my name is", p.Name)
}
