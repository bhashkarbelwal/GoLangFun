package main

import "fmt"

func main(){
	// 6) calling pyramid of stars
	fmt.Println()
	starpyramid(5)



	// 7) calling pyramid of numbers
	fmt.Println()
	numberpyramid(5)

}

// 6) pyramid of stars

func starpyramid( n int){
	for i:=0;i<n;i++{
		for j:=0;j<n-i;j++{
			fmt.Print(" ")
		}
		for k:=0;k<=i;k++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}


// 7) pyramid of numbers

func numberpyramid( n int){
	for i:=0;i<n;i++{
		for j:=0;j<n-i;j++{
			fmt.Print(" ")
		}
		for k:=0;k<=i;k++{
			fmt.Print(i+1," ")
		}
		fmt.Println()
	}
}


