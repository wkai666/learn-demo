package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 8, 10, 23, 35, 48, 59, 62, 73}
	fmt.Printf("s before remove: %+v\n", s)
	sr := remove(s, 2)
	fmt.Printf("s after remove: %+v\n", sr)
	sx := removeX(sr, 3)
	fmt.Printf("s after removeX: %+v\n", sx)
	sy := removeY(sx, 8)
	fmt.Printf("s after removeY: %+v\n", sy)
	sz := removeZ(sy, 48)
	fmt.Printf("s after removeZ: %+v\n", sz)
	si := removeY(sz, 23)
	fmt.Printf("s after removeI: %+v\n", si)
	sj := removeY(si, 62)
	fmt.Printf("s after removeJ: %+v\n", sj)

	fmt.Printf("s is: %+v\n", s)
	fmt.Println("=====================")

	sk := removeK(s, 35)
	fmt.Printf("sk is: %+v\n", sk)

	fmt.Println("=====================")
	fmt.Printf("s is: %+v\n", s)

	skk := removeK(s, 73)
	fmt.Printf("skk is: %+v\n", skk)
}

// remove 根据索引位置删除元素，删除元素后，后续位置元素前移
func remove(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

// removeX 根据索引位置删除元素，用最后一个元素覆盖删除位置的元素
func removeX(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// removeY 删除指定元素，删除元素后，后续位置元素前移 (截取法，修改原切片)
func removeY(s []int, v int) []int {
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

// removeZ 使用一个新的 slice 拷贝原切片除要删除元素之外的元素 (拷贝法，不修改原切片)
func removeZ(s []int, v int) []int {
	tmp := make([]int, 0, len(s))
	for _, vv := range s {
		if vv != v {
			tmp = append(tmp, vv)
		}
	}
	return tmp
}

// removeI 移位法修改原切片 (性能最佳)
func removeI(s []int, v int) []int {
	j := 0
	for _, vv := range s {
		if vv != v {
			s[j] = vv
			j++
		}
	}
	return s[:j]
}

func removeK(s []int, v int) []int {
	i := 0
	for _, vv := range s {
		// fmt.Printf("i is: %d, vv is: %d\n", i, vv)
		if vv == v {
			continue
		} else {
			s[i] = vv
			i++
		}
		// fmt.Printf("s is :%v\n", s)
	}
	return s[:i]
}

func removeKK(s []int, v int) []int {
	i := 0
	for _, vv := range s {
		if vv == v {
			continue
		} else {
			s[i] = vv
			i++
		}
	}
	return s[:i]
}

// removeJ 移除切片元素
func removeJ(s []int, v int) []int {
	tmp := s[:0]
	for _, vv := range s {
		if vv != v {
			tmp = append(tmp, vv)
		}
	}
	return tmp
}
