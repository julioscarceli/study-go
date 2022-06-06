package main

//Desenrolando (enumerando) uma slice
import "fmt"

func main() {

	si := []int{10, 10, 1, 2, 3, 4, 5}
	total := soma(si...)
	fmt.Println(total)
}
func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
