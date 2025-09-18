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

	var num1, num2 int

	var errCode error
	var errorFlag bool = false

	_, err1 := fmt.Scanln(&str1)
	_, err2 := fmt.Scanln(&str2)
	_, err3 := fmt.Scanln(&symbol)

	if err1 == nil {
		temp, convErr := strconv.Atoi(str1)
		if convErr != nil && !errorFlag {
			errCode = errors.New("Invalid first operand")
			errorFlag = true
		} else {
			num1 = temp
		}
	} else {
		errCode = errors.New(err1.Error())
		errorFlag = true
	}
	if err2 == nil {
		temp, convErr := strconv.Atoi(str2)
		if convErr != nil && !errorFlag {
			errCode = errors.New("Invalid second operand")
			errorFlag = true
		} else {
			num2 = temp
		}
	} else {
		errCode = errors.New(err2.Error())
		errorFlag = true
	}
	if err3 != nil {
		errCode = errors.New(err3.Error())
		errorFlag = true
	} else {
		if !validate(symbol) && !errorFlag {
			errCode = errors.New("Invalid operation")
			errorFlag = true

		} else if (str2 == "0") && (symbol == "/") && !errorFlag {
			errCode = errors.New("Division by zero")
			errorFlag = true
		}
	}

	return num1, num2, symbol, errCode
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
