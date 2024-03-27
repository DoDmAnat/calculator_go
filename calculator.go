package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
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

func isRomanNumeral(number string) bool {
	for romanNumeral := range romanNumerals {
		if number == romanNumeral {
			return true
		}
	}
	return false
}

func convertToInt(s string) int {
	num, _ := strconv.Atoi(s)
	if num >= 1 && num <= 10 {
		return num
	} else {
		panic("Числа должны быть от 1 до 10 включительно")
	}

}

func convertToArabic(romanNum string) int {
	return romanNumerals[romanNum]
}

func convertToRoman(arabicNum int) string {
	var allRoman = []struct {
		digit int
		roman string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	romanNum := ""
	for arabicNum > 0 {
		for _, pair := range allRoman {
			digit, roman := pair.digit, pair.roman
			for arabicNum >= digit {
				romanNum += roman
				arabicNum -= digit
			}
		}
	}
	return romanNum
}

func calculate(firstNum, secondNum int, operator string) int {
	switch operator {
	case "+":
		return firstNum + secondNum
	case "-":
		return firstNum - secondNum
	case "*":
		return firstNum * secondNum
	case "/":
		return firstNum / secondNum
	default:
		panic("Калькулятор не такой умный")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите арифметическое выражение:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	expression := strings.Split(input, " ")
	if len(expression) != 3 {
		panic("Cтрока не является математической операцией или " +
			"формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	firstNum, secondNum, operator := expression[0], expression[2], expression[1]
	isRomanFirst := isRomanNumeral(firstNum)
	isRomanSecond := isRomanNumeral(secondNum)
	if (isRomanFirst && !isRomanSecond) || (isRomanSecond && !isRomanFirst) {
		panic("Используются одновременно разные системы счисления")
	}
	if isRomanFirst && isRomanSecond {
		result := calculate(convertToArabic(firstNum), convertToArabic(secondNum), operator)
		if result < 1 {
			panic("В римской системе нет отрицательных чисел")
		}
		println(convertToRoman(result))
	} else {
		println(calculate(convertToInt(firstNum), convertToInt(secondNum), operator))
	}
}
