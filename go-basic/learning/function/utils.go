package function

import "fmt"

func Calculate(num1 float64, num2 float64, operator string) (result float64) {
	defer fmt.Printf("test defer \n")
	switch operator {
	case "+":
		result = num1 + num2
		break
	case "-":
		result = num1 - num2
		break
	case "x":
		result = num1 * num2
		break
	case "/":
		result = num1 / num2
		break
	default:
		result = num1 + num2
	}
	fmt.Printf("test defer before\n")
	return result
}

func Triple(n int) (result int) {
	defer func() {
		result += n
	}()
	return n + n
}
