package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nums = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var inversenums = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

func performOperation(num1, num2 int, operand string) int {
	switch operand {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		fmt.Println("Недопустимый операнд, используйте либо +, -, * или /")
		os.Exit(1)
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите расчет для выполнения: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	parts := strings.Fields(text)
	if len(parts) != 3 {
		fmt.Println("Неверный ввод, ожидаемый формат: operand1 operation operand2\n")
		os.Exit(1)
	}
	operand1, operand2, operation := parts[0], parts[2], parts[1]
	if _, ok := nums[operand1]; ok { // Если в качестве входных данных используются римские цифры
		num1, num2 := nums[operand1], nums[operand2]
		result := performOperation(num1, num2, operation)

		fmt.Println("Результат: ", inversenums[result])
	} else { // Если в качестве входных данных используются арабские цифры
		num1, _ := strconv.Atoi(operand1)
		num2, _ := strconv.Atoi(operand2)
		result := performOperation(num1, num2, operation)

		fmt.Println("Результат: ", result)
	}
}
