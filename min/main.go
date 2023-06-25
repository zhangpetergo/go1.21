package main

import "fmt"

func main() {
	x, y := 10, 20
	m := min(x) // x
	fmt.Println(m)

	m = max(x, y) // y
	fmt.Println(m)

	m = max(x, y, 100) // 100
	fmt.Println(m)

	c := max(1.0, 2.0, 10)
	fmt.Println(c)
}
