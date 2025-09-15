package main

import "fmt"
const age = 30
var name = "mike"
func main (){
	const lastName = "smith"

	const(
		port = 5000
		host = "localhost"
	)

	fmt.Println(age, name, lastName)
	fmt.Println(port, host)
}