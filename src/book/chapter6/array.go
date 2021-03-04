package main

import "fmt"

func main() {

	var x = [5]float64{ 98, 93, 77, 82, 83 }

	var total float64 = 0
	for i := 1; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	var sum float64 = 0
	for _, value := range x {
		sum += value
	}
	fmt.Println(total / float64(len(x)))

}