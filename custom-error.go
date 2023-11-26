package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		if validationError, ok := err.(*validationError); ok {
			fmt.Println(validationError.Message)
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println(notFoundError.Message)
		} else {
			fmt.Println(err.Error())
		}
	}
}

func SaveData(id string, data interface{}) error {
	if id == "" {
		return &validationError{"Error: Validation"}
	}

	if id != "aditya" {
		return &notFoundError{"Error: Not found"}
	}

	return nil
}
