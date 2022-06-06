package main

import "fmt"

const (
	_ = 1996 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)

}
