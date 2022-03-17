package main

import "fmt"

func main() {
	a := map[int]int{
		0: 0, 1: 1,
	}
	for i, _ := range a {
		a[i+2] = i + 2
	}

	b := []int{0}
	for i, _ := range b {
		b = append(b, i+1)
	}

	fmt.Println("=====>", a)
	fmt.Println("=====>", b)
}
