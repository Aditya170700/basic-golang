package main

import "fmt"

func main() {
	// init array 1
	var names [3]string
	names[0] = "Aditya"
	names[1] = "Ricki"
	names[2] = "Julianto"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// names[3] = "Alex" -> error karena panjangnya ngga bisa ditambah

	// init array 2
	values := [3]int{
		90,
		100,
		90,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// values[3] = 80 -> error karena panjangnya ngga bisa ditambah

	// init array 3
	alphabets := [...]string{
		"A",
		"B",
		"C",
	}

	fmt.Println(alphabets)
	fmt.Println(alphabets[0])
	fmt.Println(alphabets[1])
	fmt.Println(alphabets[2])

	// alphabets[3] = "Wkwkwk" -> error karena panjangnya ngga bisa ditambah

	// hitung panjang array
	fmt.Println(len(alphabets))
}
