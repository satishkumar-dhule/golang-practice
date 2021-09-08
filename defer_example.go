package main

import (
	"fmt"
)

func cleanup(s string) {

	fmt.Println("Cleaning up:", s)
}

func work() {

	defer cleanup("A")
	defer cleanup("B")

	fmt.Println("Working")
}

func main() {

	work()
}
