package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите исходное число")
	var num int
	fmt.Scan(&num)

	fmt.Println("Введите номер бита, который вы хотите изменить")
	var n int
	fmt.Scan(&n)

	fmt.Println("Введите значение, на которое вы хотите заменить данный бит")
	var c int
	fmt.Scan(&c)

	var result int
	if c == 1 {
		result = num | (1 << (n - 1))
	} else {
		result = num &^ (1 << (n - 1))
	}

	fmt.Println("Результат:")
	fmt.Println(result)
}
