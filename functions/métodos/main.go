package main

// método = funçao anexada a um tipo
import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia")
}
func main() {
	marcos := pessoa{"Marcos", 30}
	marcos.oibomdia()

}
