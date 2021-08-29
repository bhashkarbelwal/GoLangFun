package main

import "fmt"

func main(){
	var r1, r2 int
	fmt.Printf("Enter the range of numbers := ")
	fmt.Scanf("%d %d",&r1,&r2)
	var sum = 0
	fmt.Printf("Number between %d and %d, divisible by 9 := \n",r1,r2)
	for i:=r1;i<=r2;{
		if i%9 == 0{
			fmt.Print(i," ")
			sum += i
			i += 9
		} else {
			i++
		}
	}
	fmt.Printf("\nThe Sum := %d",sum)
}
