package main

import "fmt"

func main() {

ok()
}


func ok() {


	//Op- 1 1 2 3 4 5 6 7  8 9
var 	arr1 [10]int
arr1[0] =1
arr1[1] =1
arr1[2] =2
arr1[3] =3
arr1[4] =4
arr1[5] =5
arr1[6] =6
arr1[7] =7
arr1[8] =8
arr1[9] =9
fmt.Println(arr1)

//
apple := [10]int{1,1,2,3,4,5,6,7,8,9}
fmt.Println(apple)
//op=1
//   1
//   2
//  "
//  "
//  9
 number := [10]int{1,1,2,3,4,5,6,7,8,9}
fmt.Println("elements in array is : ")

 for i:=0; i<len(number);i++{
 	fmt.Println(number[i])
 }

	var 	arr2 [10]int
 arr2[0]=1
 fmt.Printf("d",arr2[0])
 for i:=1; i<len(arr2); i++{

	 fmt.Printf("%d",arr2[i])


 }



}