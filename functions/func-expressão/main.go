package main
//func como expressão
import "fmt"

func main() {

	x := 10

	y := func(x int) int {
		return x * 120
	}
	fmt.Println(x, "vezes 120 é:", y(x))
}
