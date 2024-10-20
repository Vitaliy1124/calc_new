package main

import (
	"errors"
	"fmt"
	"strings"
)


var romanToArabicMap = map[rune]int{
	'I': 1,
	"II": 2,
	'V': 5,
	'X': 10,
	
}


var arabicToRomanMap = []struct {
	Value  int
	Symbol string
}{
	{10, "X"},
	{8,"IIX"}
	{9, "IX"},
	{7, "VII"}
	{6, "VI"}
	{5, "V"},
	{4, "IV"},
	{3, "III"}
	{2, "II"}
	{1, "I"},
}


func romanToArabic(roman string) (int, error) {
	roman = strings.ToUpper(roman) 
	total := 0
	prevValue := 0

	
	for _, char := range reverseString(roman) {
		value, exists := romanToArabicMap[char]
		if !exists {
			return 0, errors.New("недопустимо: " + string(char))
		}

		
		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}

	return total, nil
}

func arabicToRoman(num int) (string, error) {
	if num <= 0 {
		return "", errors.New("panic")
	}

	roman := ""

	
	for _, entry := range arabicToRomanMap {
		for num >= entry.Value {
			roman += entry.Symbol 
			num -= entry.Value   
		}
	}

	return roman, nil
}


func reverseString(s string) string {
	runes := []rune(s) 
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func arabicCalculator(a int, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль невозможно")
		}
		return a / b, nil
	default:
		return 0, errors.New("недопуститмо: " + operator)
	}
}

func romanCalculator(a string, b string, operator string) (string, error) {
	arabicA, err := romanToArabic(a)
	if err != nil {
		return "", err
	}

	arabicB, err := romanToArabic(b)
	if err != nil {
		return "", err
	}

	result, err := arabicCalculator(arabicA, arabicB, operator)
	if err != nil {
		return "", err
	}

	romanResult, err := arabicToRoman(result)
	if err != nil {
		return "", err
	}

	return romanResult, nil
}

func main() {
	var mode string
	var operator string
	var input1, input2 string

	fmt.Println("Выберите режим: 'arabic' 'roman' :")
	fmt.Scanln(&mode)


	if mode == "arabic" {
		var a, b int

		fmt.Println("Введите первое арабское число:")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println("Ошибка ввода числа:", err)
			return
		}

		fmt.Println("Введите оператор (+, -, *, /):")
		fmt.Scanln(&operator)

		fmt.Println("Введите второе арабское число:")
		_, err = fmt.Scanln(&b)
		if err != nil {
			fmt.Println("Ошибка ввода числа:", err)
			return
		}

		
		result, err := arabicCalculator(a, b, operator)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		fmt.Printf("Результат: %d\n", result)

	} else if mode == "roman" {
		

		fmt.Println("Введите первое римское число:")
		fmt.Scanln(&input1)

		fmt.Println("Введите оператор (+, -, *, /):")
		fmt.Scanln(&operator)

		fmt.Println("Введите второе римское число:")
		fmt.Scanln(&input2)

		
		result, err := romanCalculator(input1, input2, operator)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		fmt.Printf("Результат в римских цифрах: %s\n", result)

	} else {
		fmt.Println("Недопустимый режим. Выберите 'arabic' или 'roman'.")
	}