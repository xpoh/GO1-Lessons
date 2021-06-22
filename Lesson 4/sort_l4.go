package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertInSlice(sl []float64, val float64, pos int) []float64 {
	sl = append(sl, 0)
	n := len(sl)
	for i := n - 1; i > pos; i-- {
		sl[i] = sl[i-1]
	}
	sl[pos] = val
	return sl
}

func sortInsertion(A []float64) []float64 {
	for i := 1; i < len(A); i++ {
		x := A[i]
		j := i
		for (j > 0) && (A[j-1] > x) {
			A[j] = A[j-1]
			j--
		}
		A[j] = x
	}
	return A
}

func main() {
	var nums []float64
	var sortCount int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите набор чисел для сортировки (в одну строку разделенный пробелом):")
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	for _, snum := range s {
		num, err := strconv.ParseFloat(snum, 64)
		if err == nil {
			nums = append(nums, num)
		} else {
			panic("not number")
		}
	}
	out := make([]float64, 1)
	out[0] = nums[0]
	sortCount = 1
	for i := 1; i < len(nums); i++ {
		fInserted := false
		for j := 0; j < sortCount; j++ {
			if nums[i] <= out[j] {
				out = insertInSlice(out, nums[i], j)
				sortCount++
				fInserted = true
				break
			}
		}
		if !fInserted {
			out = insertInSlice(out, nums[i], sortCount)
			sortCount++
		}
	}
	fmt.Print("Отсортированный массив:")
	fmt.Println(out[0:sortCount])

	fmt.Print("Отсортированный массив by sortInsertion:")
	fmt.Println(sortInsertion(nums))
}
