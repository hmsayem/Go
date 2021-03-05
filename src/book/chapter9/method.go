package main
import "fmt"

type Student struct{
	name string
	grades []int
	age int
}
func (s *Student) getAge() int {
	return s.age
}
func (s *Student) setAge(age int) {
	s.age = age
}
func (s *Student) getAvgGrade() float64{
	sum := 0
	for _,grade := range s.grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.grades))
}
func main(){
	s1 := Student{"Sayem", []int{90,78,87,50}, 20}
	fmt.Println(s1.getAge())
	s1.setAge(21)
	fmt.Println(s1.getAge())
	fmt.Println(s1.getAvgGrade())
}