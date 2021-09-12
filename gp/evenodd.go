package main

import "fmt"

func main() {
	var asize,i,evenCount,oddCount int
	fmt.Println("enter the array size:")
	fmt.Scanln(&asize)
	array:= make([]int,asize)
	fmt.Println("enter the array values:")
	for i=0;i<asize;i++{
		fmt.Scanln(&array[i])
	}
	evenCount=0
	oddCount=0
	for _, a := range array{
		if a%2 ==0{
			evenCount++
	} else{
		oddCount++
	}
	}
	fmt.Println("the even numbers are:", evenCount)
	fmt.Println("the odd numbers are:", oddCount)


}
