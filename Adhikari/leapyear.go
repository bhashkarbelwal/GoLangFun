package main

import "fmt"

func leap_year(year int) {

	if year%4==0 && year%100!=0 || year%400==0{
		fmt.Println("year is leap year",year)
	}else {
		fmt.Println("Year is not leap year", year)
	}
}
func main(){
	leap_year(1600)
}




