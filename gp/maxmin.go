package main

import "fmt"

func main() {
	var max,min,asize,i int
	fmt.Println("enter the array size:")
	fmt.Scanln(&asize)
	array:= make([]int,asize)
	fmt.Println("enter the array values")
	for i=0;i<asize;i++{
	fmt.Scanln(&array[i])
	}
	max = array[0]
	min = array[0]
	for i=0;i<asize;i++{
		if max<array[i]{
			max=array[i]

		}
		if min>array[i]{
			min=array[i]

		}
	}
	fmt.Println("the maximum number is:",max)

	fmt.Println("the minimum number is:",min)

}
