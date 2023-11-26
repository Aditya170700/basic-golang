package main

import "fmt"

type MyLocation struct {
	City, Province, Country string
}

func main() {
	location1 := new(MyLocation)
	location2 := location1

	location2.City = "Jogja"

	fmt.Println(location1)
	fmt.Println(location2)
}
