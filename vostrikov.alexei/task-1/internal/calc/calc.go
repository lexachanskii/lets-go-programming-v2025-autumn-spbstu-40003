package calc

import (
	"fmt"
	"os"
)

func Calculate() {

	var num1 int
	var num2 int
	var symbol string

	if _, err := fmt.Scanln(&num1); err != nil {
		fmt.Println("Invalid first operand")
		os.Exit(1)
	}
	if _, err := fmt.Scanln(&num2); err != nil {
		fmt.Println("Invalid second operand")
		os.Exit(1)
	}
	if _, err := fmt.Scanln(&symbol); err != nil {
		fmt.Println("Problem in symbol")
		os.Exit(1)
	} else {
		if !validate(symbol) {
			fmt.Println("Invalid operation")
			os.Exit(1)
		} else if (num2 == 0) && (symbol == "/") {
			fmt.Println("Division by zero")
			os.Exit(1)
		}
	}

	switch symbol {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(float32(num1) / float32(num2))
	}
}

func validate(str string) bool {

	var correct [4]string = [4]string{"-", "+", "*", "/"}

	for i := 0; i < len(correct); i++ {
		if str == correct[i] {
			return true
		}
	}
	return false
}
