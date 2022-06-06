package main

//Anonymous self-executing functions → Funções anônimas auto-executáveis
import "fmt"

func main() {
	x := 150

	func(x int) {
		fmt.Println(x, "vezes 87654 é: ")
		fmt.Println(x * 87654)
	}(x)
}
