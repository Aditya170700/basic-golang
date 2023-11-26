package main

import "fmt"

func main() {
	runApplication(true)
}

func runApplication(error bool) {
	// tetap dieksekusi diakhir program walaupun error
	defer logging()

	fmt.Println("Run Application")

	if error {
		// program akan berhenti
		panic("Error guys")
	}
}

func logging() {
	// mengambil message panic
	message := recover()
	fmt.Println("Error :", message)
}
