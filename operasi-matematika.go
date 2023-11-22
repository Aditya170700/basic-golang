package main

import "fmt"

func main() {
	fmt.Println(1 + 1) // tambah (2)
	fmt.Println(1 - 1) // kurang (0)
	fmt.Println(1 * 1) // kali (1)
	fmt.Println(1 / 1) // bagi (1)
	fmt.Println(1 % 1) // sisa hasil bagi (0)

	// augmented assigmnetn
	a := 10

	a += 10
	fmt.Println(a) // a = a + 10
	a -= 10
	fmt.Println(a) // a = a - 10
	a *= 10
	fmt.Println(a) // a = a * 10
	a /= 10
	fmt.Println(a) // a = a / 10
	a %= 10
	fmt.Println(a) // a = a % 10
}
