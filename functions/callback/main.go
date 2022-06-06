package main

import "fmt"

func main() {

	x := Somenteimpares(soma, []int{7, 7, 7, 7, 2}...)
	fmt.Println(x)

}
func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v //n = n+v
	}
	return n
}
func Somenteimpares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}
