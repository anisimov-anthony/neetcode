package stack

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, 2)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			if token == "+" {
				result = a + b
			} else if token == "-" {
				result = a - b
			} else if token == "*" {
				result = a * b
			} else { // "/"
				result = a / b
			}

			stack = append(stack, result)

		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
