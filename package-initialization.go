package main

import (
	"basic-golang/database"
	"fmt"
)

func main() {
	fmt.Println(database.GetConnection())
}
