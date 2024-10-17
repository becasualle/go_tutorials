package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "hello"
	printMe(printValue)

	numerator, denominator := 4, 3
	result, remainder, err := intDivision(numerator, denominator)

	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("Результат деления: %v", result)
	// } else {
	// 	fmt.Printf("Результат деления: %v, остаток: %v", result, remainder)
	// }

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("Результат деления: %v", result)
	default:
		fmt.Printf("Результат деления: %v, остаток: %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("Результат деления: %v", result)
	default:
		fmt.Printf("Результат деления: %v, остаток: %v", result, remainder)
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("целое нельзя делить на ноль")
		return 0, 0, err
	}

	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}
