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

	if _, err := fmt.Scanln(&departments); err != nil {
		return fmt.Errorf("error while reading departments count: %w", err)
	}

	for range departments { // Go 1.22+: диапазон 0..departments-1
		if _, err := fmt.Scanln(&employee); err != nil {
			return fmt.Errorf("error while reading employee count: %w", err)
		}

		if err := tempControl(employee); err != nil {
			return fmt.Errorf("error in temperature control: %w", err)
		}
	}

	return nil
}

func tempControl(employee int) error {
	var (
		symbol string
		temp   string
	)

	lower := 15
	higher := 30
	var broken bool

	for range employee {
		if _, err := fmt.Scanln(&symbol, &temp); err != nil {
			return fmt.Errorf("error while reading temperature line: %w", err)
		}

		if broken {
			fmt.Println(-1)

			continue
		}

		tempInt, isHigher, err := validateTemp(symbol, temp)
		if err != nil {
			return fmt.Errorf("error while validating temperature: %w", err)
		}

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

			continue
		}

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

	return nil
}

func validateTemp(symbol, temp string) (int, bool, error) {
	value, err := strconv.Atoi(temp)
	if err != nil {
		return 0, false, fmt.Errorf("error while converting string to int: %w", err)
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
