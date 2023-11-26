package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Division(10, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}
}

func Division(pembilang int, penyebut int) (int, error) {
	if penyebut == 0 {
		return 0, errors.New("Error: penyebut 0")
	} else {
		return pembilang / penyebut, nil
	}
}
