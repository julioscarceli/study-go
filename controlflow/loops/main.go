package main

import "fmt"

func main() {

	for x := 0; x <= 4; x++ {
		fmt.Println(x)

	}

	// The most basic type, with a single condition
	x := 0
	for x < 4 {
		fmt.Println(x)
		x++
	}

	//for break
	anoquenasci := 1996
	anoatual := 2022

	for {
		if anoquenasci == anoatual {
			break
		}
		fmt.Println(anoquenasci)
		anoquenasci++

	}

	// for continue
	for x := 0; x < 5; x++ {
		if x == 3 {
			continue
		}
		fmt.Println(x)
	}

}
