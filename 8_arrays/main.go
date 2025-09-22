package main

import (
	"fmt"
)

func main(){

	//zerod vals
	var nums [4]int
	fmt.Println(len(nums))

	nums[0]=1
	fmt.Println(nums)


	var vals[4]bool
	vals[0]=true
	fmt.Println(vals)

	var names [3]string
	names[2]="golang"
	fmt.Println(names)

//single line declaration
	number :=[3]int{1,2,3}
	fmt.Println(number)


	// 2d arrays
	nums2d := [2][3]int{{1,2},{4,5,6}}
	fmt.Println(nums2d)

	//fixed size, this is predictable
	//memory optimization
	//constant time access
}
