package main

import (
	"fmt"
	"os"
	"strconv"
)

type Operation func(int, int) int

func readInput() (string, string, string) {
	var a, opp, b string
	fmt.Fscan(os.Stdin, &a)
	fmt.Fscan(os.Stdin, &opp)
	fmt.Fscan(os.Stdin, &b)
	return a, opp, b
}

func convertToArabic(num string) int {
	a := make(map[string]int)
	a["I"] = 1
	a["II"] = 2
	a["III"] = 3
	a["IV"] = 4
	a["V"] = 5
	a["VI"] = 6
	a["VII"] = 7
	a["VII"] = 8
	a["IX"] = 9
	a["X"] = 10
	result, ok := a[num]
	if ok {
		return result
	} else {
		panic("Слишком большие числа")
	}
}

func convertToRoman(num int) string {
	roman := ""
	for num > 0 {
		switch {
		case num >= 100:
			roman += "C"
			num -= 100
		case num >= 90:
			roman += "XC"
			num -= 90
		case num >= 50:
			roman += "L"
			num -= 50
		case num >= 40:
			roman += "XL"
			num -= 40
		case num >= 10:
			roman += "X"
			num -= 10
		case num >= 9:
			roman += "IX"
			num -= 9
		case num >= 5:
			roman += "V"
			num -= 5
		case num >= 4:
			roman += "IV"
			num -= 4
		case num >= 1:
			roman += "I"
			num -= 1
		}
	}
	return roman
}



func calcResult(a int, opp string, b int) int {
	opps := map[string]Operation{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	if a > 10 || b > 10 {
		panic("Слишком большие числа")
	}
	result, ok := opps[opp]
	if ok {
		return result(a, b)
	} else {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
}

func main() {
	str_a, opp, str_b := readInput()
	a, err1 := strconv.Atoi(str_a)
	b, err2 := strconv.Atoi(str_b)
	if err1 != nil && err2 != nil {
		a := convertToArabic(str_a)
		b := convertToArabic(str_b)
		result := calcResult(int(a), opp, int(b))
		if result < 0 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел")
		}
		fmt.Println(convertToRoman(result))
	

	} else if err1 != nil || err2 != nil {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	} else {
		fmt.Println(calcResult(int(a), opp, int(b)))
	}
}
