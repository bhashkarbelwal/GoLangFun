package main

import "fmt"

func main(){
	var limit int
	fmt.Println("Enter the number of fibonacci terms you want to print")
	fmt.Scanln(&limit)
	fibo(limit)
}

func fibo(n int){
	fmt.Print("0 1 ")
	n1 := 0
	n2 := 1
	for i:=2;i<n;i++{
		n3 := n1 + n2
		fmt.Print(n3," ")
		n1 = n2
		n2 = n3


	}
}
