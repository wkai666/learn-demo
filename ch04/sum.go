package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 4))
	fmt.Println(sum(1, 2, 4, 5, 6))

	// 切片作为参数进行传递
	values := []int{3, 5, 6, 8, 9}
	fmt.Println(sum(values...))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
