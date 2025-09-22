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


	//type switch
	whoAmI := func (i interface{})  {
		switch t:=i.(type) {
		case int:
			fmt.Println("its a int")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("unknown type", t)
		}
	}
	whoAmI("haanji")
	whoAmI(52)
	whoAmI(true)
	whoAmI(55.286)
}
