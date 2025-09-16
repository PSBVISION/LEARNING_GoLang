package main

import "fmt"
// for -> only loop construct in golang
func main (){
	// while loop
	i:=1
	for i<5{
		fmt.Println(i)
		i=i+1
	}
	//infinite loop
	// for{
	// 	fmt.Println("hello")
	// }

	//for loop

	for j:=0;j<5;j++{
		fmt.Println(j)
	}

	//range
	for i:= range 3 {
		println(i)
	}
}