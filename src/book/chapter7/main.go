package main
import "fmt"

func add(x, y int) int{
	return x + y
}
func divide(x, y int) (int, int){
	return x / y, x % y
}
func addSub(x, y int) (a, s int){
	a = x + y
	s = x - y
	return
}
func testDefer(){
	defer fmt.Println("After")
	fmt.Println("Before")
}
func test(do func(int) int) int{
	return do(13)
}
func returnFunc(x string ) func(){
	return func(){
		fmt.Println(x)
	}
}
func main() {

	x := 50
	y := 7

	sum1 := add(x,y)
	fmt.Println(sum1)
	do  := add
	sum2 :=do(x,y)
	fmt.Println(sum2)

	d, r  := divide(x, y)
	fmt.Println(d, r)

	a, s := addSub(x, y)
	fmt.Println(a,s)

	testDefer()

	getNeg := func(x int) int {
		if x > 0{
			return x * -1
		}
		return x
	}
	v1 := getNeg(100)
	fmt.Println(v1)

	getPos := func(x int) int{
		if x < 0{
			return x * -1
		}
		return x
	}
	v2 := getPos(-100)
	fmt.Println(v2)

	v3 := test(getNeg)
	fmt.Println(v3)

	v4 := test(getPos)
	fmt.Println(v4)

	sq := func(x int) int{
		return x * x
	}(11)
	fmt.Println(sq)

	func(){
		fmt.Println("Hello World")
	}()

	c1 := returnFunc("Hello")
	c2 := returnFunc("World")
	c1()
	c2()
}