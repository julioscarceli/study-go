package main

import "fmt"

//Demonstre o resto da divisão por 4 de todos os números entre 10 e 20

func main() {
	for x := 10; x <= 20; x++ {
		fmt.Println(x % 4)
	}

}
