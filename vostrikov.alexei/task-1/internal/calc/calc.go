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

	_, err1 := fmt.Scanln(&str1)
	_, err2 := fmt.Scanln(&str2)
	_, err3 := fmt.Scanln(&symbol)

	if err1 == nil {
		temp, convErr := strconv.Atoi(str1)
		if convErr != nil {
			errCode = errors.New("Invalid first operand")
		} else {
			num1 = temp
		}
	} else {
		errCode = errors.New(err1.Error())
	}
	if err2 == nil {
		temp, convErr := strconv.Atoi(str2)
		if convErr != nil {
			errCode = errors.New("Invalid second operand")
		} else {
			num2 = temp
		}
	} else {
		errCode = errors.New(err2.Error())
	}
	if err3 != nil {
		errCode = errors.New(err3.Error())
	} else {
		if !validate(symbol) {
			errCode = errors.New("Invalid operation")
		} else if (str2 == "0") && (symbol == "/") {
			errCode = errors.New("Division by zero")
		}
	}
	return num1, num2, symbol, errCode
}

func Calculate() (float32, error) {
	num1, num2, symbol, errCode := input()

	var res float32

	if errCode != nil {
		return res, errCode
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
