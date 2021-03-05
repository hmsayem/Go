package main

import (
	"fmt"
	"math/rand"
)

func changeValue(ptr *int){
	* ptr = rand.Int()
}
func main(){

	x := 2
	changeValue(&x)
	fmt.Println(x)

}
