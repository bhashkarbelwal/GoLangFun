package main

import "fmt"

func bigestno(a , b, c int){

	if a>b && a>c{
		fmt.Println("a is greater no",a)
	}else if b>c {
		fmt.Println("b is greater no",b)

	}else{
		fmt.Println("c is greater no",c)
	}
}
func main(){
	bigestno(150,120,25)

}