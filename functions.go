package main

import (
	"fmt"
)

func add(a int, b int) int {

	return a + b
}

func devmod(a int, b int) (int, int) {

	return a / b, a % b
}

func doubleAt(array []int, ix int) {

	array[ix] *= 2
}

func main() {

	val := add(1, 2)
	fmt.Println(val)

	d, m := devmod(30, 4)
	fmt.Println(d, m)

	a := []int{1, 2, 3, 4, 5, 6}
	doubleAt(a, 2)

	fmt.Println(a)

}
