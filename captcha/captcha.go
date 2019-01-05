package captcha

import "fmt"

func Do(pattern int, leftO int, operator int, rightO int) string {
	result := ""

	if rightO > 9 {
		return "righ operand should not greather than 9"
	}

	numberWord := [9]string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	operatorWord := [4]string{
		"+",
		"-",
		"*",
		"/",
	}

	if pattern == 1 {
		result = fmt.Sprintf("%v %v %v", leftO, operatorWord[operator-1], numberWord[rightO-1])
	}

	return result
}
