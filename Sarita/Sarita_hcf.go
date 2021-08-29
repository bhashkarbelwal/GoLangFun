package main

import "fmt"

func main(){
	var n1,n2,i int
	fmt.Printf("Enter 1st Number := ")
	fmt.Scanf("%d",&n1)
	fmt.Printf("Enter 2nd Number := ")
	fmt.Scanf("%d",&n2)
	if n1 > n2{
		i = n2
	} else {
		i = n2
	}
	for ;i>0;i--{
		if n1%i==0 && n2%i==0 {
			fmt.Printf("HCF of %d and %d is := %d",n1,n2,i)
			break
		}
	}
}
