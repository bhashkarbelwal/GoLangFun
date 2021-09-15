package main

import "fmt"

func fabonaciseries(number int) {
	a := 0
	b := 1
	c := 0

	fmt.Print("Enter the number")
	for i := 0; i <= number; i++ {
		if i == 0 {
			fmt.Print(" ", a)
			continue
		} else if i == 2 {
			fmt.Print(" ", b)
			continue
		}
		c = a + b
		a = b
		b = c
		fmt.Print(" ", c)

	}
}

func main(){
		fabonaciseries(20)
	}






