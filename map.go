package main

import "fmt"

func main() {
	// cara init map 1
	var person1 map[string]string = make(map[string]string)
	person1["name"] = "Aditya"
	person1["address"] = "Jogja"

	// cara init map 1
	person2 := map[string]string{
		"name":    "Ricki",
		"address": "Jogja",
		"wrong":   "Ups",
	}

	fmt.Println(person1)
	fmt.Println(person1["name"])
	fmt.Println(person1["address"])

	fmt.Println(person2)
	fmt.Println(person2["name"])
	fmt.Println(person2["address"])

	// hapus data map
	delete(person2, "wrong")

	fmt.Println(person2)
}
