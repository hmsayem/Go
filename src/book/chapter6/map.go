package main

import "fmt"

func main() {

	var mp1 = map[string]int{
		"H":  1,
		"He": 2,
		"Li": 3,
		"Be": 4,
	}
	mp1["B"] = 5
	mp1["C"] = 6
	fmt.Println(mp1)

	delete(mp1, "Be")
	fmt.Println(mp1)

	val, ok := mp1["Li"]
	fmt.Println(val, ok)

	var mp2 = make(map[string]int)
	fmt.Println(mp2)
}