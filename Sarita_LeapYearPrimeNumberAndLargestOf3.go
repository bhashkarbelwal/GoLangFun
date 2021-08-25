package main

import "fmt"

func main(){
	// 2) calling Isprime function
	fmt.Println()
	var p int
	p = Isprime(97)
	if p==1 {
		fmt.Println("prime number")
	} else {
		fmt.Println("Not a prime number")
	}


	// 3) Leap Year
	fmt.Println()
	yr := IsLeapYear(2000)
	if yr {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not a Leap Year")
	}


	// 4) largest of 3
	fmt.Println()
	largestof3(56, 89, 71)


}


// 2) To print number is prime or not

func Isprime(n int) int {
	for i:=2;i<n/2;i++ {
		if n%i==0{
			return 1
		}

	}
	return 0
}

// To check the given year is leap year or not

func IsLeapYear(y int) bool{
	if y%400==0 || (y%4==0 && y%100!=0){
		return true
	}
	return false
}


// 4) Function to find largest of 3 numbers

func largestof3(n1 int, n2 int, n3 int) {
	if n1>=n2 {
		if n3 > n1{
			fmt.Println(n3, "is largest ammong all the 3 numbers")
		} else {
			fmt.Println(n1," is largest among all the 3 numbers")
		}
	} else {
		if n3 > n2{
			fmt.Println(n3, "is largest ammong all the 3 numbers")
		} else {
			fmt.Println(n2," is largest among all the 3 numbers")
		}
	}
}
