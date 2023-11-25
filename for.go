package main

import "fmt"

func main() {
	// cara 1
	counter := 1
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	// cara 2
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// cara 3
	names := []string{"Aditya", "Ricki", "Julianto"}
	// dengan index / key
	for i, name := range names {
		fmt.Println(i+1, name)
	}

	// tanpa index / key
	for _, name := range names {
		fmt.Println(name)
	}
}
