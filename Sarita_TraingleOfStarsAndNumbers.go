package main
import "fmt"
func main (){
// 1) calling triangle of stars
startrainglepattern(5)

	// 5) Right angled triangle of numbers
	fmt.Println()
	rightangletriangleofnumber(4)
}



//  1) To print the triangle of stars

func startrainglepattern(n int) {
	for i:=0;i<n;i++ {
		for j:=0;j<i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}


// 5) Right angled triangle of numbers

func rightangletriangleofnumber(n int){
	num := 1
	for i:=0;i<n;i++{
		for j:=0;j<=i;j++{
			fmt.Print(num, " ")
			num++
		}
		fmt.Println()
	}
}