package main

import "fmt"
func main(){
	fmt.Println("Enter the number")
	var num int
	fmt.Scanln(&num)
	Ispalindrome(num)
}

func Ispalindrome(n int){
	temp := n
	var newn int = 0
	for n:=n;n>0;n=n/10 {
		no := n%10
		newn += no*no*no
	}
	if temp == newn{
		fmt.Printf(" %d is Armstrong number\n", temp)
	} else {
		fmt.Printf(" %d is not Armstrong number\n\n", temp)
	}
}
