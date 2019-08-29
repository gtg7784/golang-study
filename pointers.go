package main

import "fmt"

func main() {
	x := 0
	changeVal(x)
	fmt.Println("x =", x)

	changeXValNow(&x)
	fmt.Println("x =", x)

	fmt.Println("Memory Address for x =", &x)
}

func changeVal(x int) {
	x = 2
}

func changeXValNow(x *int) {
	*x = 2
}
