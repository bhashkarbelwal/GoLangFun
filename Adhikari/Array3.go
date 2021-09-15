package main
import "fmt"
func main(){
	var num int
	fmt.Println("enter the number")
	fmt.Scanln(&num)
	fmt.Println("Original Number := ",num)
	rev := 0
	for i:=num;i>0;i=i/10{
		d := i%10
		rev = rev*10+d
	}
	fmt.Println("Reversed Number := ",rev)
}

