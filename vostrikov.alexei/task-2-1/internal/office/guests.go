package guests

import (
	"errors"
	"fmt"
	"strconv"
)

func ClimateControl() error {
	var (
		departments int
		employee    int
	)

	_, err := fmt.Scanln(&departments)
	if err != nil {
		return fmt.Errorf("error while reading departments count: %w", err)
	}

	for department := 0; department < departments; department++ {
		_, err := fmt.Scanln(&employee)
		if err != nil {
			return fmt.Errorf("error while reading employee count: %w", err)
		}
		err = tempControl(employee)
		if err != nil {
			return fmt.Errorf("error in temprature control: %w", err)
		}
	}
	return nil
}

func tempControl(employee int) error {
	var (
		temp   string
		symbol string
		lower  int  = 15
		higher int  = 30
		broken bool = false
	)
	for current := 0; current < employee; current++ {
		_, err := fmt.Scanln(&symbol, &temp)
		if err != nil {
			return fmt.Errorf("error while reading temprature line : %w", err)
		}

		if broken {
			fmt.Println(-1)
			continue
		}

		tempInt, isHigher, err := validateTemp(symbol, temp)
		if err != nil {
			return fmt.Errorf("error while validating temprature : %w", err)
		} else {
			if isHigher {
				if tempInt > higher {
					broken = true
					fmt.Println(-1)
					continue
				}
				if tempInt > lower {
					lower = tempInt
				}
				fmt.Println(lower)
			} else {
				if tempInt < lower {
					broken = true
					fmt.Println(-1)
					continue
				}
				if tempInt < higher {
					higher = tempInt
				}
				fmt.Println(lower)
			}
		}
	}
	return nil
}

func validateTemp(symbol string, temp string) (int, bool, error) {

	value, err := strconv.Atoi(temp)
	if err != nil {
		return 0, false, fmt.Errorf("error while converting string to int : %w", err)
	}
	switch symbol {
	case "<=":
		return value, false, nil
	case ">=":
		return value, true, nil
	default:
		return 0, false, errors.New("could not match symbol - " + symbol)

	}
}
