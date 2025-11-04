package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var f float64 = 12.8
	var i int = int(f)
	fmt.Println(i)

	//type inference
	var w int
	j := w
	fmt.Println(j)

}

var n, c, e bool
var t, y = true, false

const New = 2.65
