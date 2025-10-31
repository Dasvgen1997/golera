package main

import "fmt"

func main() {
	isEvenNum := isEven(11)

	fmt.Println(isEvenNum, " is even num")

	maxNum := max(1, 5)

	fmt.Println(maxNum, " max num")

	gradeName := getGrade(74)

	fmt.Println(gradeName, " is grade name")

}

func isEven(num int) bool {
	var isEven bool = (num % 2) == 0
	return isEven
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getGrade(score int) string {
	if score >= 90 {
		return "Отлично"
	} else if score >= 70 {
		return "Хорошо"
	} else if score >= 50 {
		return "Удовлетворительно"
	} else {
		return "Неудовлетворительно"
	}
}
