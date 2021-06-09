package main

import (
	"fmt"
)

func isSimple(num int) bool {
	var result = true

	for i := 2; i < num-1; i++ {
		if num%i == 0 {
			result = false
		}
	}
	return result
}

func main() {
	var num int
	var count int
	fmt.Print("Введите максимальное значение простого числа:")
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	for i := 0; i < num; i++ {
		if isSimple(i) {
			count++
			fmt.Print(i, "\t")
			if count%8 == 0 {
				fmt.Println("")
			}
		}
	}
	fmt.Printf("\nНайдено %d простых чисел\n", count)
	fmt.Println("Сложность алгоритма O(n^2)")
}
