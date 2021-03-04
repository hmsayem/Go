package main

import "fmt"
func main() {

	var x = [5]int{1, 2, 3, 4, 5}
	var s1 = x[1:3]
	fmt.Println(s1)
	s1 = s1[1:3]
	fmt.Println(s1)

	var s2 = []int{1, 2, 3, 4, 5}
	fmt.Println(cap(s2[:3]))
	s2 = append(s2, 6, 7)
	fmt.Println(s2)

	var s3 = make([]int, 5)
	fmt.Println(s3)
	fmt.Printf("%T", s3)
}