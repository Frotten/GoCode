package main

import (
	"fmt"
)

func Add(Arr ...int) (sum int) {
	for _, value := range Arr {
		sum += value
	}
	return sum
}
func main() {
	Arr := []int{2, 3, 6, 9, 5}
	fmt.Println("Sum :", Add(Arr...))
}
