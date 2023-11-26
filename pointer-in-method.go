package main

import "fmt"

type Man struct {
	Name string
}

func (m Man) GetFullNameWithoutPointer() {
	m.Name = "Mr. " + m.Name
}

func (m *Man) GetFullNameWithPointer() {
	m.Name = "Mr. " + m.Name
}

func main() {
	aditya := Man{Name: "Aditya"}
	aditya.GetFullNameWithoutPointer()
	// name tidak berubah
	fmt.Println(aditya)

	ricki := Man{Name: "Ricki"}
	ricki.GetFullNameWithPointer()
	// name akan berubah "Mr. Ricki"
	fmt.Println(ricki)
}
