package main

import "fmt"

func main() {

	var aInt8 int = 3
	var bInt16 int16 = 3
	var cInt32 int32 = 3
	var dInt64 int64 = 3
	var eString string = "test"
	var fFloat32 float32 = 32.3
	var gFloat64 float64 = 32.3
	var hinterface interface{} = 32.3

	fmt.Println(aInt8)
	fmt.Println(bInt16)
	fmt.Println(cInt32)
	fmt.Println(dInt64)
	fmt.Println(eString)
	fmt.Println(fFloat32)
	fmt.Println(gFloat64)
	fmt.Println(hinterface)
}
