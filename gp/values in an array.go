package main

import "fmt"

func main() {
		var asize, i int
		fmt.Println("enter array size:")
		fmt.Scanln(&asize)
		array:= make([]int,asize)
		fmt.Scanln(&array)
		fmt.Println("enter the array values:")
		for i=0;i<asize;i++{
			fmt.Scanln(&array[i])
		}
		fmt.Println("the array is:",array)


	}


