package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func calc(a, b, o string) (float32, error) {
	var err error

	fa, err := strconv.ParseFloat(a, 32)
	if err != nil {
		return 0, err
	}

	fb, err := strconv.ParseFloat(b, 32)
	if err != nil {
		return 0, err
	}

	switch o {
	case "+":
		return float32(fa + fb), nil
	case "-":
		return float32(fa - fb), nil
	case "*":
		return float32(fa * fb), nil
	case "/":
		return float32(fa / fb), nil
	default:
		return 0, errors.New("операция неизвестна")
	}
}

func main() {
	var str string
	var oper [4]string
	oper[0] = "+"
	oper[1] = "-"
	oper[2] = "*"
	oper[3] = "/"

	fmt.Print("Введите выражение (допускается только 2 числа и операции +-*/): ")
	fmt.Scanln(&str)

	for _, s := range oper {
		ss := strings.Split(str, s)
		if len(ss) == 2 {
			result, err := calc(ss[0], ss[1], s)
			if err == nil {
				fmt.Printf("Ответ: %f\n", result)
			} else {
				fmt.Printf("Ошибка: %s\n", err)
			}
			break
		}
	}
}
