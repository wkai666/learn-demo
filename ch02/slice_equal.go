package main

import "fmt"

func main() {
	a := []string{"a", "c", "d", "e"}
	b := []string{"a", "e", "d", "c"}
	c := []string{"a", "c", "d", "e"}

	fmt.Printf("a equal b is: %t\n", equal(a, b))
	fmt.Printf("a equal c is: %t\n", equal(a, c))
}

func equal(x, y []string) bool {
	// 切片长度的比较
	if len(x) != len(y) {
		return false
	}

	// 切片中每个元素的值及索引的比较
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}
