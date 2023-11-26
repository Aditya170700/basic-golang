package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Aditya", spamFilter)
}

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}
