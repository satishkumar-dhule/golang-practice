package main

import (
	"fmt"
)

func cleanup(s string) {

	fmt.Println("Cleaning up:", s)
}

func work() {

	defer cleanup("A")

	fmt.Println("Working")
}

func main() {

	work()
}
