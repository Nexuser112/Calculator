package main

import (
	"bufio"
	"errors"
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
	0: "0", 1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI",
	7: "VII", 8: "VIII", 9: "IX", 10: "X", 11: "XI", 12: "XII", 13: "XIII",
	14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
	21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX",
	30: "XXX", 31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV", 36: "XXXVI",
	37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL", 41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV",
	46: "XLVI", 47: "XLVII", 48: "XLVIII",
	49: "XLIX", 50: "L", 51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV", 56: "LVI", 57: "LVII", 58: "LVIII",
	59: "LIX", 60: "LX", 61: "LXI", 62: "LXII",
	63: "LXIII", 64: "LXIV", 65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX", 71: "LXXI",
	72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX",
	80: "LXXX", 81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV",
	86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC", 91: "XCI", 92: "XCII", 93: "XCIII",
	94: "XCIV", 95: "XCV", 96: "XCVI", 97: "XCVII",
	98: "XCVIII", 99: "XCIX", 100: "C",
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
		err := errors.New("Недопустимый операнд, используйте либо +, -, * или /")
		fmt.Print(err)
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
		err1 := errors.New("Неверный ввод, ожидаемый формат: operand1 operation operand2\n")
		fmt.Print(err1)
		os.Exit(1)
	}
	operand1, operand2, operation := parts[0], parts[2], parts[1]

	if _, ok := nums[operand1]; ok { // Если в качестве входных данных используются римские цифры
		num1, num2 := nums[operand1], nums[operand2]
		result := performOperation(num1, num2, operation)
		if result <= 0 {
			err2 := errors.New("Результат должен быть больше нуля")
			fmt.Print(err2)
			os.Exit(1)
		}
		if num1 < 1 || num1 > 10 {
			err3 := errors.New("Числа должны находится в диапазоне от 1 до 10")
			fmt.Print(err3)
			os.Exit(1)
		}
		if num2 < 1 || num2 > 10 {
			err3 := errors.New("Числа должны находится в диапазоне от 1 до 10")
			fmt.Print(err3)
			os.Exit(1)
		}
		fmt.Println("Результат: ", inversenums[result])
	} else { // Если в качестве входных данных используются арабские цифры
		num1, _ := strconv.Atoi(operand1)
		num2, _ := strconv.Atoi(operand2)
		result := performOperation(num1, num2, operation)
		if num1 < 1 || num1 > 10 {
			err3 := errors.New("Числа должны находится в диапазоне от 1 до 10")
			fmt.Print(err3)
			os.Exit(1)
		}
		if num2 < 1 || num2 > 10 {
			err3 := errors.New("Числа должны находится в диапазоне от 1 до 10")
			fmt.Print(err3)
			os.Exit(1)
		}

		fmt.Println("Результат: ", result)
	}
}
