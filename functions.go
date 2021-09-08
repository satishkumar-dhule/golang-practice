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

func double(a int) {

	a *= 2
}

func doublePtr(a *int) {

	*a *= 2
}

func main() {

	val := add(1, 2)
	fmt.Println(val)

	d, m := devmod(30, 4)
	fmt.Println(d, m)

	a := []int{1, 2, 3, 4, 5, 6}
	doubleAt(a, 2)

	fmt.Println(a)

	i := 10

	double(i)
	fmt.Println(i) //no change in value as i awas passed by value

	doublePtr(&i)
	fmt.Println(i) //change in value as i was passed by ref

}
