package main

import "fmt"

func main() {
	var C int
	fmt.Scan(&C)
	fmt.Println("Hello, World!")
	const (
		c0 = iota
		c1
		c2
		c3
	)
	// goto LOOP
	var A, B int = 10, 20
	fmt.Println(C)
	nums := [10]int{1, 2, 3, 6, 5, 4, 9, 8, 7, 10}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	fmt.Println(c0, c1, c2, c3)
	fmt.Println(A, B)
}
