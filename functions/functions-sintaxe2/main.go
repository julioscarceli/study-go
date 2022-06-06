package main

//funçoes com multiplos retornos
import "fmt"

func main() {
	total, quantos := soma(20, 20, 20, 20, 20)
	fmt.Println(total, quantos)
}
func soma(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)

}
