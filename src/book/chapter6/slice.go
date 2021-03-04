package main

import "fmt"
func main() {

	var x  = [5]int{1,2,3,4,5}
	var s  = x[1:3]
	fmt.Println(s)
	s = x[1: 4]
	fmt.Println(s)
	s = s[1:3]
	fmt.Println(s)
}