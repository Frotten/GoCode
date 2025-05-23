package sort1

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

type TempFunc func(int, int) int

func AnyFunc(a, b int, LL TempFunc) int {
	return LL(a, b)
}

func main() {
	fmt.Println("Add:", AnyFunc(1, 2, Add))
	fmt.Println("Sub:", AnyFunc(1, 2, Sub))
}
