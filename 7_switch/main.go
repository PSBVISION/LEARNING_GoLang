package main

import (
	"fmt"
	"time"
)

// for -> only loop construct in golang
func main(){
	//simple switch
	i:=5
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}


	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: 
		fmt.Println("it's weekend")
	default: 
		fmt.Println("it's a weekday")
	}
}
