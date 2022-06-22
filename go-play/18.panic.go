package main

import "fmt"

func divide(x, y int) {
	if y == 0 {
		panic("You can not divide a number by Zero")
	}
	fmt.Printf("%d / %d = %d\n", x, y, x/y)
}

func main() {
	divide(10, 5)
	divide(20, 0)
	divide(30, 10)
}