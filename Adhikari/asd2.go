package main

import "fmt"

func kalki() {

	for i := 0; i <= 5; i++ {
		for j := 0; j <= i; j++ {


				fmt.Print("*")
			}
					fmt.Println()
		}
	for  k := 5; k >= 0; k-- {
		fmt.Print("*")
	}
	fmt.Println()

}
func main(){
	kalki()
}