package calc

import (
	"errors"
	"fmt"
	"strconv"
)

func input() (int, int, string, error) {

	var str1 string
	var str2 string
	var symbol string

	var err_code error
	var error_flag bool = false

	if _, err := fmt.Scanln(&str1); err == nil {

		if _, err := strconv.Atoi(str1); err != nil && !error_flag {
			err_code = errors.New("invalid first operand")
			error_flag = true
		}
	} else {
		err_code = errors.New(err.Error())
		error_flag = true
	}

	if _, err := fmt.Scanln(&str2); err == nil {

		if _, err := strconv.Atoi(str2); err != nil && !error_flag {
			err_code = errors.New("invalid second operand")
			error_flag = true
		}
	} else {
		err_code = errors.New(err.Error())
		error_flag = true
	}

	if _, err := fmt.Scanln(&symbol); err != nil {

		err_code = errors.New("problem in symbol")
		error_flag = true

	} else {
		if !validate(symbol) && !error_flag {

			err_code = errors.New("invalid operation")
			error_flag = true

		} else if (str2 == "0") && (symbol == "/") && !error_flag {

			err_code = errors.New("division by zero")
			error_flag = true

		}
	}

	num1, _ := strconv.Atoi(str1)
	num2, _ := strconv.Atoi(str2)

	return num1, num2, symbol, err_code
}

func Calculate() (float32, error) {

	num1, num2, symbol, err_code := input()

	var res float32 = 0.0

	if err_code != nil {
		return res, err_code
	}

	switch symbol {
	case "+":
		res = float32(num1 + num2)
	case "-":
		res = float32(num1 - num2)
	case "*":
		res = float32(num1 * num2)
	case "/":
		res = float32(num1) / float32(num2)
	}

	return res, nil
}

func validate(str string) bool {

	var correct = [4]string{"-", "+", "*", "/"}

	for i := 0; i < len(correct); i++ {
		if str == correct[i] {
			return true
		}
	}
	return false
}
