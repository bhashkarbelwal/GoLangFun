package main

import "fmt"

func main() {
	var asize,i,j int
	fmt.Println("enter array size:")
	fmt.Scanln(&asize)
	array := make([]int, asize)
	rarray := make([]int, asize)
	fmt.Println("enter the array values:")
	for i=0;i<asize;i++{
		fmt.Scanln(&array[i])
	}
	j=0
	for i=asize-1;i>=0;i--{
		rarray[j] =array[i]
		j++
	}
	fmt.Println("reverse array is:",rarray)

}
