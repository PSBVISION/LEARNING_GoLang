package main

import "fmt"

// for -> only loop construct in golang
func main() {
	age := 20
	if age > 18 {
		fmt.Println("you can vote")
	} else {
		fmt.Println("you can't vote")
	}

	if age>=18{
		fmt.Println("you can vote")
	}else if age>=12{
		fmt.Println("person is a teenager")
	}

	var role = "admin"
	var hasPermissions = true
	if role == "admin" && hasPermissions {
		fmt.Println("hello admin")
	}

	if age:=15; age>=18{
		fmt.Println("user is adult", age)
	}else if age>=12{
		fmt.Println("user is teenager", age)
	}else{
		fmt.Println("user is child", age)
	}
	// go deosn't have ternary operator
}
