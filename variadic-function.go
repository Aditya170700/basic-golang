package main

import "fmt"

func main() {
	// parameter angka biasa
	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	// parameter slice
	total = sumAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(total)
}

func sumAll(numbers ...int) (total int) {
	for _, number := range numbers {
		total += number
	}

	return total
}
