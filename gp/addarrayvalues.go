package main

import "fmt"

func main() {
	var asize, i, asum int
	fmt.Println("enter array size:")
	fmt.Scanln(&asize)
	fmt.Println("enter the array values:")
	array:= make([]int,asize)
	fmt.Scanln(&array)
	for i=0;i<asize;i++{
		fmt.Scanln(&array[i])

	}
	fmt.Println(" the array values are:",array)
	asum = 0
	for i=0;i<asize;i++{
		asum=asum+array[i]
	}
	fmt.Println("sum of values of an array is",asum)


}
