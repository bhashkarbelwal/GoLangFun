package main
import "fmt"
func main(){
	var num int
	fmt.Println("enter the number")
	fmt.Scanln(&num)
	IsPalindrome(num)
}

func IsPalindrome(n int){
	temp := n
	rev := 0
	for i:=n;i>0;i=i/10{
		d := i%10
		rev = rev*10+d
	}
	if temp == rev {
		fmt.Printf("%d is palindrome number\n", temp)
	} else {
		fmt.Printf("%d is not a palindrome number\n",temp)
	}
}