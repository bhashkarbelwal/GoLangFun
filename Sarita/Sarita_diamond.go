package main

import "fmt"

func main(){
	fmt.Println("Enter the size for your diamond")
	var n int
	fmt.Scanln(&n)
	stardiamond(n)
}
func stardiamond( n int){
	for i:=0;i<=n;i++{
		for j:=0;j<n-i;j++{
			fmt.Print(" ")
		}
		for k:=1;k<=2*i+1;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i:=1;i<=n;i++{
		for j:=0;j<i;j++{
			fmt.Print(" ")
		}
		for k:=2*(n-i)+1;k>=1;k--{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
