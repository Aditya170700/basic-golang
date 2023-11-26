package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// contoh 1
	address1 := Address{"Jogja", "DIY", "Indonesia"}

	// pass by value
	//address2 := address1

	// (pointer) pass by reference
	address2 := &address1

	address2.City = "Sleman"

	fmt.Println(address1)
	fmt.Println(address2)

	// contoh 2
	var address3 Address = Address{"Kulonprogo", "DIY", "Indonesia"}
	var address4 *Address = &address3

	address4.City = "Bantul"

	fmt.Println(address3)
	fmt.Println(address4)
}
