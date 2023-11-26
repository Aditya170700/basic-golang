package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	// cara 1
	var customer1 Customer
	customer1.Name = "Aditya"
	customer1.Address = "Jogja"
	customer1.Age = 23

	fmt.Println(customer1)
	fmt.Println(customer1.Name)
	fmt.Println(customer1.Address)
	fmt.Println(customer1.Age)

	// cara 2
	customer2 := Customer{
		Name:    "Ricki",
		Address: "Jogja",
		Age:     23,
	}

	fmt.Println(customer2)
	fmt.Println(customer2.Name)
	fmt.Println(customer2.Address)
	fmt.Println(customer2.Age)

	// cara 3
	customer3 := Customer{"Julianto", "Jogja", 23}

	fmt.Println(customer3)
	fmt.Println(customer3.Name)
	fmt.Println(customer3.Address)
	fmt.Println(customer3.Age)
}
