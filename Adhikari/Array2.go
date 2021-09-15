package main

import "fmt"

func main() {
	numb := []int{2,5,8}
	fmt.Print("Result :", sum(numb))
}



func sum(array []int) int {
result := 0
for _, v := range array {
result += v
}
return result
}
