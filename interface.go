package main

import (
	"fmt"
)

type HasName interface {
	GetName() string
}

type Student struct {
	Name string
}

type Teacher struct {
	Name string
}

func main() {
	student1 := Student{Name: "Aditya"}
	SayHello(student1)

	teacher1 := Teacher{Name: "Ricki"}
	SayHello(teacher1)
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// implementasi interface di struct student
func (s Student) GetName() string {
	return s.Name
}

// implementasi interface di struct teacher
func (t Teacher) GetName() string {
	return t.Name
}
