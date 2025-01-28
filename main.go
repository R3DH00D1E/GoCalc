package main

import "fmt"

func main() {
	var firstNum, secondNum, result float32
	const (
		slozhenie  = 1
		vichitanie = 2
		umnozhenie = 3
		delenie    = 4
		ostatok    = 5
	)
	var operation int
	fmt.Println("Добро пожаловать в Калькулятор!!!")
	fmt.Println("Введите два числа через пробел:")
	fmt.Scan(&firstNum, &secondNum)
	fmt.Printf("Введённые числа: %g и %g\n", firstNum, secondNum)
	fmt.Println("Выберите операцию:")
	fmt.Println("1. Сложение")
	fmt.Println("2. Вычитание")
	fmt.Println("3. Умножение")
	fmt.Println("4. Деление")
	fmt.Println("5. Остаток от деления")
	fmt.Scan(&operation)
	switch operation {
	case slozhenie:
		result = firstNum + secondNum
	case vichitanie:
		result = firstNum - secondNum
	case umnozhenie:
		result = firstNum * secondNum
	case delenie:
		if secondNum == 0 {
			fmt.Println("Делить на ноль нельзя!")
			break
		} else {
			result = firstNum / secondNum
		}
	case ostatok:
		result = float32(int(firstNum) % int(secondNum))
	}
	fmt.Printf("Результат: %g\n", result)
}
