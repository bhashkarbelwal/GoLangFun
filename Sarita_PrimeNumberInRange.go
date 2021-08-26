package main

import "fmt"

func main(){
	var lower, upper int
	fmt.Println("Enter the lower limit")
	fmt.Scanln(&lower)
	fmt.Println("Enter the upper limit")
	fmt.Scanln(&upper)
	for j:=lower;j<=upper;j++{
		ans := Isprime(j)
		if ans==1 {
			fmt.Print(j," ")
		}
	}

}

func Isprime(n int) int {
	if n==1{
		return 0
	}
	for i:=2;i<n/2;i++ {
		if n%i==0{
			return 0
		}
	}
	return 1
}