package main

import "fmt"

const (
	_ = iota
	a = 1 << iota
	b
)

const (
	c = iota
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
