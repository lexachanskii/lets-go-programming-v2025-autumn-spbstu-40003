package main

import (
	"fmt"

	table "github.com/lexachanskii/task-2-2/internal"
)

func main() {
	dish, err := table.Table()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(dish)
}
