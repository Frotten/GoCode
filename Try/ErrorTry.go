package main

import (
	"errors"
	"fmt"
)

func Fa(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("zero division")
	} else {
		return a / b, nil
	}
}
func main() {
	result, err := Fa(1, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
