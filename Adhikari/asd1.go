package main

import "fmt"

func bigfoot(){

	for i:=0; i<5;i++{
		for j:=1; j<i; j++{
fmt.Printf("%d ",i+j)

}
		fmt.Println()
	}

}
func main(){
	bigfoot()
}
