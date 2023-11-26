package main

import "fmt"

type MyAddress struct {
	City, Province, Country string
}

func main() {
	var address1 MyAddress = MyAddress{"Kulonprogo", "DIY", "Indonesia"}
	var address2 *MyAddress = &address1

	address2.City = "Bantul"

	fmt.Println(address1)
	fmt.Println(address2)

	//address2 = &MyAddress{"Sleman", "DIY", "Indonesia"}
	*address2 = MyAddress{"Sleman", "DIY", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
