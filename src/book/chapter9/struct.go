package main

import (
	"fmt"
	"math/rand"
)

type Point struct{
	x int32
	y int32
}
type Circle struct{
	radius float64
	center Point
}
func changePoint(ptr *Point){
	ptr.x = rand.Int31()
	ptr.y = rand.Int31()
}
func main(){

	p := Point{1,2}
	fmt.Println(p)
	changePoint(&p)
	fmt.Println(p)

	c := Circle{10, p}
	fmt.Println(c)

}
