package main

import "fmt"

func main() {
	s := []int{4, 6, 8, 9, 12, 34, 56, 89}

	fmt.Printf("s is %+v\n", s)
	s = insert(s, 5, 27)
	fmt.Printf("after insert, s is: %+v\n", s)
}

// insert 向切片中 index 的位置插入一个元素 value，原先的元素位置后移
func insert(s []int, index, value int) []int {
	if len(s)+1 > cap(s) {
		s = append(s, value)
	}
	// fmt.Printf("s len: %d, cap: %d\n", len(s), cap(s))
	// fmt.Printf("s is: %+v\n", s)

	for i := len(s) - 1; i >= 0; i-- {
		if i > index {
			s[i] = s[i-1]
		} else if i == index {
			s[i] = value
		} else {
			continue
		}
	}

	return s
}
