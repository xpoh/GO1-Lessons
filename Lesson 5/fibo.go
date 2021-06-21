package main

import (
	fmt "fmt"
)

func getFibo(n int) int {
	var a, b int = 0, 1
	var result int
	if n == 0 {
		result = 0
	} else if n == 1 {
		result = 1
	} else {
		for i := 1; i < n; i++ {
			result = a + b
			a = b
			b = result
		}
	}
	return result
}

func main() {
	var n int

	fmt.Print("Введите номер числа Фибоначи:")
	fmt.Scan(&n)

	fmt.Print("Чило Фибоначи:")
	fmt.Println(getFibo(n))
}
