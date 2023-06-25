package main

import "fmt"

func main() {
	s := []string{
		"hello",
		"world",
	}

	fmt.Println(len(s))
	fmt.Println(s)

	clear(s)

	fmt.Println(len(s))
	fmt.Println(s)

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s1))
	fmt.Println(s1)

	clear(s1)

	fmt.Println(len(s1))
	fmt.Println(s1)

	m := map[string]interface{}{
		"fuck":     "you",
		"fuck you": "two",
	}

	fmt.Println(len(m))
	fmt.Println(m)

	clear(m)

	fmt.Println(len(m))
	fmt.Println(m)

}
